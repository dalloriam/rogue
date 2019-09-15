package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestPosition_Name(t *testing.T) {
	d := &components.Position{}
	if d.Name() != components.PositionName {
		t.Error("invalid name")
	}
}
