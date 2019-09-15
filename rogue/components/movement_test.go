package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestMovement_Name(t *testing.T) {
	d := &components.Movement{}
	if d.Name() != components.MovementName {
		t.Error("invalid name")
	}
}
