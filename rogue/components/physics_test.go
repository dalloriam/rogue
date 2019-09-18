package components_test

import (
	"testing"

	"github.com/purposed/good/datastructure/stringset"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestPhysics_Name(t *testing.T) {
	d := &components.Physics{}
	if d.Name() != components.PhysicsName {
		t.Error("invalid name")
	}
}

func TestPhysics_IsBlocked(t *testing.T) {
	d := &components.Physics{BlockedBy: stringset.FromValues([]string{"wall"})}

	if !d.BlockedBy.Contains("wall") {
		t.Error("expected physics component to be blocked by wall")
		return
	}

	if d.BlockedBy.Contains("floor") {
		t.Error("did not expect physics component to be blocked by floor")
		return
	}
}
