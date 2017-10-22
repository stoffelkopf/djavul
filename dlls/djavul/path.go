package main

// #include <stdint.h>
// typedef int32_t bool32_t;
//
// static bool32_t is_valid_step(bool32_t (*__fastcall f)(int entity_num, int x, int y), int entity_num, int x, int y) {
//    return f(entity_num, x, y);
// }
//
// static int __fastcall path_make(bool32_t (__fastcall *is_valid_step)(int entity_num, int x, int y), int entity_num, int start_x, int start_y, int target_x, int target_y, int8_t *steps) {
//    int (__fastcall *f)(bool32_t (__fastcall *is_valid_step)(int entity_num, int x, int y), int entity_num, int start_x, int start_y, int target_x, int target_y, int8_t *steps) = (void *) 0x4493D4;
//    return f(is_valid_step, entity_num, start_x, start_y, target_x, target_y, steps);
// }
import "C"

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"

	"github.com/kr/pretty"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
)

// MakePath returns the shortest path as a sequence of steps for the given
// entity from start to target. The step validation function reports valid
// movements.
//
// ref: 0x4493D4

//export MakePath
func MakePath(validStep unsafe.Pointer, entityNum, startX, startY, targetX, targetY int32, steps *int8) int {
	// Go implementation.
	valid := func(x, y int) bool {
		return C.is_valid_step((*[0]byte)(validStep), C.int(entityNum), C.int(x), C.int(y)) == 1
	}
	sh := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(steps)),
		Len:  25,
		Cap:  25,
	}
	s := *(*[]Step)(unsafe.Pointer(&sh))

	// Original implementation.
	m := C.path_make((*[0]byte)(validStep), C.int(entityNum), C.int(startX), C.int(startY), C.int(targetX), C.int(targetY), (*C.int8_t)(unsafe.Pointer(&s[0])))
	fmt.Println("s orig:", s[:m])

	n := makePath(valid, int(startX), int(startY), int(targetX), int(targetY), s)
	fmt.Println("s new: ", s[:n])
	return n
}

func stepFromNodes(prev, n *Node) Step {
	pretty.Println("prev:", prev)
	pretty.Println("n:", n)
	dx := n.x - prev.x
	dy := n.y - prev.y
	fmt.Println("dx:", dx)
	fmt.Println("dy:", dy)
	switch {
	case dx == -1 && dy == -1:
		return StepN
	case dx == -1 && dy == 0:
		return StepNW
	case dx == -1 && dy == 1:
		return StepW
	case dx == 0 && dy == -1:
		return StepNE
	case dx == 0 && dy == 0:
		// no movement.
		panic("invalid step; no movement")
	case dx == 0 && dy == 1:
		return StepSW
	case dx == 1 && dy == -1:
		return StepE
	case dx == 1 && dy == 0:
		return StepSE
	case dx == 1 && dy == 1:
		return StepS
	}
	panic("unreachable")
}

type Step int8

const (
	StepNE Step = 1
	StepNW Step = 2
	StepSE Step = 3
	StepSW Step = 4
	StepN  Step = 5
	StepE  Step = 6
	StepS  Step = 7
	StepW  Step = 8
)

func makePath(validStep func(x, y int) bool, startX, startY, targetX, targetY int, steps []Step) int {
	fmt.Printf("start (%d, %d), target (%d, %d)\n", startX, startY, targetX, targetY)
	const (
		width  = 96
		height = 96
	)
	g := newGrid(width, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if validStep(x, y) {
				id := int64(x + y*width)
				g.cells[x][y] = &Node{
					id: id,
					x:  x,
					y:  y,
				}
			}
		}
	}
	// Enture target node is in grid, even if not reachable. The path finding
	// algorithm may then walk up right next to the target cell, and stop one
	// step away.
	id := int64(targetX + targetY*width)
	g.cells[targetX][targetY] = &Node{
		id: id,
		x:  int(targetX),
		y:  int(targetY),
	}

	h := func(u, v graph.Node) float64 {
		uu := node(u)
		vv := node(v)
		w := abs(uu.x - vv.x)
		h := abs(uu.y - vv.y)
		return math.Sqrt(float64(w*w + h*h))
	}
	shortest, _ := path.AStar(g.cells[startX][startY], g.cells[targetX][targetY], g, h)
	nodes, weight := shortest.To(g.cells[targetX][targetY])
	pretty.Println("nodes:", nodes)
	fmt.Println("weight:", weight)

	// Translate nodes to steps.
	n := 0
	for i := 1; i < len(nodes); i++ {
		prev := node(nodes[i-1])
		nn := node(nodes[i])
		step := stepFromNodes(prev, nn)
		// TODO: Skip last step (the one to the target cell) if not valid step.
		steps[n] = step
		n++
	}
	pretty.Println("steps:", steps)
	return n
}

type Grid struct {
	cells  [][]*Node
	width  int
	height int
}

func newGrid(width, height int) *Grid {
	cells := make([][]*Node, width)
	for x := range cells {
		cells[x] = make([]*Node, height)
	}
	g := &Grid{
		cells:  cells,
		width:  width,
		height: height,
	}
	return g
}

type Node struct {
	// id = x + y*width
	id int64
	x  int
	y  int
}

func (n *Node) ID() int64 {
	return n.id
}

func (g *Grid) Has(n graph.Node) bool {
	nn := node(n)
	return g.cells[nn.x][nn.y] != nil
}

func (g *Grid) Nodes() []graph.Node {
	var nodes []graph.Node
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			n := g.cells[x][y]
			if n != nil {
				nodes = append(nodes, n)
			}
		}
	}
	return nodes
}

func (g *Grid) From(n graph.Node) []graph.Node {
	var nodes []graph.Node
	nn := node(n)
	for x := nn.x - 1; x <= nn.x+1; x++ {
		for y := nn.y - 1; y <= nn.y+1; y++ {
			if x < 0 || x >= g.width {
				continue
			}
			if y < 0 || y >= g.height {
				continue
			}
			if x == nn.x && y == nn.y {
				continue
			}
			if g.cells[x][y] != nil {
				nodes = append(nodes, g.cells[x][y])
			}
		}
	}
	return nodes
}

func (g *Grid) HasEdgeBetween(u, v graph.Node) bool {
	uu := node(u)
	vv := node(v)
	if g.cells[uu.x][uu.y] == nil {
		return false
	}
	if g.cells[vv.x][vv.y] == nil {
		return false
	}
	if abs(uu.x-vv.x) > 1 || abs(uu.y-vv.y) > 1 {
		return false
	}
	return !(uu.x == vv.x && uu.y == vv.y)
}

func (g *Grid) Edge(u, v graph.Node) graph.Edge {
	if !g.HasEdgeBetween(u, v) {
		return nil
	}
	e := &Edge{
		F: u,
		T: v,
	}
	return e
}

type Edge struct {
	F graph.Node
	T graph.Node
}

func (e *Edge) From() graph.Node {
	return e.F
}

func (e *Edge) To() graph.Node {
	return e.T
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func node(n graph.Node) *Node {
	if n, ok := n.(*Node); ok {
		return n
	}
	panic(fmt.Errorf("invalid node type; expected *path.Node, got %T", n))
}
