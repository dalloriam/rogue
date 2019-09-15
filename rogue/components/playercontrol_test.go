package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestPlayerControl_Name(t *testing.T) {
	d := &components.PlayerControl{}
	if d.Name() != components.PlayerControlName {
		t.Error("invalid name")
	}
}
