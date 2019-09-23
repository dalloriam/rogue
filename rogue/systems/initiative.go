package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/components"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

var minTurnDelta = 100 * time.Millisecond // TODO: Make configurable.

type InitiativeSystem struct {
	playerWentFirst bool
	turnCount       uint64

	timeOfLastTurn time.Time
}

// NewInitiativeSystem initializes and returns an initiative system.
func NewInitiativeSystem() *InitiativeSystem {
	return &InitiativeSystem{
		playerWentFirst: false,
		turnCount:       0,

		timeOfLastTurn: time.Now(),
	}
}

func (s *InitiativeSystem) ShouldTrack(obj object.GameObject) bool {
	return obj.HasComponent(components.ControlName)
}

func (s *InitiativeSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	if time.Since(s.timeOfLastTurn) < minTurnDelta {
		// No initiative possible -- not enough time.
		return nil
	}

	sawAPlayer := false
	aiPlayed := false
	for _, obj := range objects {
		if obj.HasComponent(components.InitiativeName) {
			continue
		}

		if obj.HasComponent(components.PlayerName) {
			obj.AddComponents(&components.Initiative{}) // Player always has initiative
			sawAPlayer = true
		} else if s.playerWentFirst {
			obj.AddComponents(&components.Initiative{})
			aiPlayed = true
		}
	}

	if sawAPlayer {
		s.playerWentFirst = true
	}

	if aiPlayed {
		s.playerWentFirst = false
	}

	s.turnCount++
	s.timeOfLastTurn = time.Now()

	return nil
}
