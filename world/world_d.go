//+build djavul

package world

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/sanctuary/djavul/engine"
	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/internal/proto"
	"github.com/sanctuary/djavul/lighting"
	"github.com/sanctuary/djavul/scrollrt"
)

// drawTopArchesUpperScreen draws arches on the upper screen, with added
// transparency.
//
// ref: 0x463060
func drawTopArchesUpperScreen(dstBuf unsafe.Pointer) {
	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// drawBottomArchesUpperScreen draws arches on the upper screen, with added
// transparency.
//
// ref: 0x46468D
func drawBottomArchesUpperScreen(dstBuf unsafe.Pointer, drawMasks *uint32) {
	// TODO: Figure out how to handle drawMasks.
	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// ref: 0x4652C5
func drawUpperScreen(dstBuf unsafe.Pointer) {
	if *scrollrt.CelTransparencyActive != 0 {
		switch *scrollrt.ArchDrawType {
		case 0:
			drawTopArchesUpperScreen(dstBuf)
			return
		case 1:
			solid := gendung.Solid_0x10_0x20_0x40_FromPieceID[*scrollrt.LevelPieceID]
			if solid == 1 || solid == 3 {
				// if sw wall
				//
				//     /
				//    /
				drawBottomArchesUpperScreen(dstBuf, &TileDrawMasks[63])
				return
			}
		case 2:
			solid := gendung.Solid_0x10_0x20_0x40_FromPieceID[*scrollrt.LevelPieceID]
			if solid == 2 || solid == 3 {
				// if se wall
				//
				//     \
				//      \
				drawBottomArchesUpperScreen(dstBuf, &TileDrawMasks[31])
				return
			}
		}
	}

	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// drawTopArchesLowerScreen draws arches on the lower screen, with added
// transparency.
//
// ref: 0x465F38
func drawTopArchesLowerScreen(dstBuf unsafe.Pointer) {
	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// drawBottomArchesLowerScreen draws arches on the lower screen, with added
// transparency.
//
// ref: 0x467949
func drawBottomArchesLowerScreen(dstBuf unsafe.Pointer, drawMasks *uint32) {
	// TODO: Figure out how to handle drawMasks.
	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// ref: 0x46886B
func drawLowerScreen(dstBuf unsafe.Pointer) {
	if *scrollrt.CelTransparencyActive != 0 {
		switch *scrollrt.ArchDrawType {
		case 0:
			drawTopArchesLowerScreen(dstBuf)
			return
		case 1:
			solid := gendung.Solid_0x10_0x20_0x40_FromPieceID[*scrollrt.LevelPieceID]
			if solid == 1 || solid == 3 {
				// if sw wall
				//
				//     /
				//    /
				drawBottomArchesLowerScreen(dstBuf, &TileDrawMasks[63])
				return
			}
		case 2:
			solid := gendung.Solid_0x10_0x20_0x40_FromPieceID[*scrollrt.LevelPieceID]
			if solid == 2 || solid == 3 {
				// if se wall
				//
				//    \
				//     \
				drawBottomArchesLowerScreen(dstBuf, &TileDrawMasks[31])
				return
			}
		}
	}

	x, y := engine.CalcXY(dstBuf)
	var frameNum int
	if 0 < *scrollrt.LightTableIndex && *scrollrt.LightTableIndex < int32(*lighting.Max) {
		if *scrollrt.LevelCelBlock&0x8000 != 0 {
			frameNum = int(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][0]) - 1
			// TODO: Figure out how to handle light index.
			//frameContent = speed_cels + scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[*scrollrt.LevelCelBlock&0xFFF][scrollrt.LightTableIndex]
		} else {
			frameNum = int(*scrollrt.LevelCelBlock&0xFFF) - 1
			// TODO: Figure out how to handle light index.
			//lightEntry := &LightTable[256*scrollrt.LightTableIndex]
		}
	} else {
		block := *scrollrt.LevelCelBlock
		if block&0x8000 != 0 {
			block = uint32(scrollrt.SpeedCelFrameNumFromLightIndexFrameNum[block&0xFFF][0]) + block&0xF000
		}
		frameNum = int(block&0xFFF) - 1
	}
	if err := proto.SendDrawImage(celPathFromDType(), x, y, frameNum); err != nil {
		log.Fatalf("%+v", err)
	}
}

// ### [ Helper functions ] ####################################################

// celPathFromDType returns the level CEL path for the active dungeon type.
func celPathFromDType() string {
	switch *gendung.DType {
	case gendung.Tristram:
		return "levels/towndata/town.cel"
	case gendung.Cathedral:
		return "levels/l1data/l1.cel"
	case gendung.Catacombs:
		return "levels/l2data/l2.cel"
	case gendung.Caves:
		return "levels/l3data/l3.cel"
	case gendung.Hell:
		return "levels/l4data/l4.cel"
	default:
		panic(fmt.Errorf("unknown dungeon type %d", uint(*gendung.DType)))
	}
}
