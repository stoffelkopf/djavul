package proto

// EngineEvent is a Diablo 1 game engine event.
type EngineEvent interface {
	// isEngineEvent ensures that only engine events can be assigned to the
	// EngineEvent interface.
	isEngineEvent()
}
