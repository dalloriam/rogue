package cartography_test

import (
	"testing"

	"github.com/dalloriam/rogue/rogue/cartography"
)

func TestNewMap(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m := cartography.NewMap(i, j)

			if (m.SizeY() != j && i != 0) || m.SizeX() != i {
				t.Error("invalid size reported")
				return
			}

			if actual := len(m); actual != i {
				t.Errorf("invalid x length. expected %d, got %d", i, actual)
				return
			}

			for x := 0; x < i; x++ {
				if actual := len(m[x]); actual != j {
					t.Errorf("invalid y length. expected %d, got %d", j, actual)
					return
				}
			}
		}
	}
}
