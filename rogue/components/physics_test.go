package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestPhysics_Name(t *testing.T) {
	d := &components.Physics{}
	if d.Name() != components.PhysicsName {
		t.Error("invalid name")
	}
}

func TestPhysics_IsBlocked(t *testing.T) {
	d := &components.Physics{BlockedBy: []string{"wall"}}

	if !d.IsBlocked("wall") {
		t.Error("expected physics component to be blocked by wall")
		return
	}

	if d.IsBlocked("floor") {
		t.Error("did not expect physics component to be blocked by floor")
		return
	}
}
