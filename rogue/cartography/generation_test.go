package cartography_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/dalloriam/rogue/rogue/cartography"
)

type mockTemplate struct{}

func (m *mockTemplate) Generate(source *rand.Rand) cartography.Map {
	lvlMap := cartography.NewMap(1, 1)
	lvlMap.Set(cartography.Tile{
		X:       0,
		Y:       0,
		Char:    'w',
		Type:    fmt.Sprintf("%d", source.Int()),
		FgColor: nil,
		BgColor: nil,
	})
	return lvlMap
}

func TestLevelManager_Generation(t *testing.T) {
	m := cartography.NewLevelManager("test.txt", 1234)
	_, ok := m.GetLevel("test_level")
	if ok {
		t.Error("level should not exist")
		return
	}

	m.AddLevel("test_level", &mockTemplate{})
	if lvl, ok := m.GetLevel("test_level"); ok {
		if lvl.SizeX() != 1 || lvl.SizeY() != 1 {
			t.Error("invalid map dimensions")
			return
		}

		if tile := lvl.At(0, 0); tile.Char != 'w' {
			t.Error("invalid tile character")
		}
	} else {
		t.Error("level should exist")
	}
}

func TestLevelManager_Seeding(t *testing.T) {
	m1 := cartography.NewLevelManager("test.txt", 1234)
	m2 := cartography.NewLevelManager("test2.txt", 1234)
	m3 := cartography.NewLevelManager("test2.txt", 4321)

	l1 := m1.AddLevel("test_level", &mockTemplate{})
	l2 := m2.AddLevel("other_level", &mockTemplate{})
	l3 := m3.AddLevel("yet_other_level", &mockTemplate{})

	t1 := l1.At(0, 0)
	t2 := l2.At(0, 0)
	t3 := l3.At(0, 0)

	if (t1.Type != t2.Type) || (t1.Type == t3.Type) || (t2.Type == t3.Type) {
		t.Error("invalid seeding")
	}
}
