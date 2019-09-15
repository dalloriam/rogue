package components_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"
)

func TestDrawable_Name(t *testing.T) {
	d := &components.Drawable{}
	if d.Name() != components.DrawableName {
		t.Error("invalid name")
	}
}
