package object_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/components"

	"github.com/dalloriam/rogue/rogue/object"
)

func TestNewObject(t *testing.T) {
	obj1 := object.New(&components.PlayerControl{}, &components.Position{
		X: 0,
		Y: 0,
	})

	if obj1.ID() != uint64(1) {
		t.Error("invalid ID")
		return
	}

	if !obj1.HasComponent(components.PlayerControlName) || !obj1.HasComponent(components.PositionName) {
		t.Error("components not initialized properly")
		return
	}

	if obj1.HasComponent(components.DrawableName) {
		t.Error("object has too many components")
		return
	}

	obj2 := object.New()
	if obj1.ID() == obj2.ID() {
		t.Error("IDs not incremented properly")
		return
	}
}

func TestBaseObject_AddComponents(t *testing.T) {
	obj1 := object.New()
	if obj1.HasComponent(components.PlayerControlName) {
		t.Error("object has component before adding")
		return
	}

	// Test single add.
	obj1.AddComponents(&components.PlayerControl{})
	if !obj1.HasComponent(components.PlayerControlName) {
		t.Error("component not added properly")
		return
	}

	// Test Multi-add
	obj1.AddComponents(&components.Drawable{}, &components.Physics{})
	if !obj1.HasComponent(components.DrawableName) || !obj1.HasComponent(components.PhysicsName) {
		t.Error("components not added properly")
		return
	}
}

func TestBaseObject_RemoveComponent(t *testing.T) {
	obj1 := object.New(&components.Drawable{})
	if !obj1.HasComponent(components.DrawableName) {
		t.Error("component not added")
		return
	}

	obj1.RemoveComponent(components.DrawableName)
	if obj1.HasComponent(components.DrawableName) {
		t.Error("component not removed properly")
		return
	}
}

func TestBaseObject_GetComponent(t *testing.T) {
	playerControl := &components.PlayerControl{}
	pos := &components.Position{
		X: 123,
		Y: 14,
	}
	obj1 := object.New(playerControl, pos)

	if p2 := obj1.GetComponent(components.PositionName).(*components.Position); pos != p2 {
		t.Error("wrong component returned")
		return
	}

	if ctrl2 := obj1.GetComponent(components.PlayerControlName).(*components.PlayerControl); ctrl2 != playerControl {
		t.Error("wrong component returned")
		return
	}
}
