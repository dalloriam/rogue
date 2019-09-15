package gameplay

import "go.uber.org/atomic"

// TurnClock returns the current game turn.
type TurnClock struct {
	currentTurn atomic.Uint64
}

// GetTurn returns the current turn.
func (t *TurnClock) GetTurn() uint64 {
	return t.currentTurn.Load()
}

// Increment increments the turn clock.
func (t *TurnClock) Increment() {
	t.currentTurn.Add(1)
}
