package player

// Class specifies a player class.
type Class uint8

//go:generate stringer -type Class

// Player classes.
const (
	Warrior Class = iota
	Rogue
	Sorceror
)
