package systems

import (
	"time"

	"github.com/dalloriam/rogue/rogue/components"

	"github.com/dalloriam/rogue/rogue/cartography"
	"github.com/dalloriam/rogue/rogue/object"
)

type InitiativeSystem struct {
}

// NewInitiativeSystem initializes and returns an initiative system.
func NewInitiativeSystem() *InitiativeSystem {
	return &InitiativeSystem{}
}

func (s *InitiativeSystem) Name() string {
	return "initiative"
}

func (s *InitiativeSystem) ShouldTrack(obj object.GameObject) bool {
	return obj.HasComponent(components.ControlName)
}

func (s *InitiativeSystem) Update(dT time.Duration, worldMap cartography.Map, objects map[uint64]object.GameObject) error {
	for _, obj := range objects {
		if obj.HasComponent(components.InitiativeName) {
			continue
		}

		// If we see one player we want to give them the initiative.
		if obj.HasComponent(components.PlayerName) {
			obj.AddComponents(&components.Initiative{})
		}
	}

	return nil
}
