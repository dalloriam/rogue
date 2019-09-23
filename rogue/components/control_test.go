package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestControl_Name(t *testing.T) {
	d := &components.Control{}
	if d.Name() != components.ControlName {
		t.Error("invalid name")
	}
}
