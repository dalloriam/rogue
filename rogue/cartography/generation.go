package cartography

import (
	"fmt"
	"math/rand"
)

// A LevelTemplate implements logic for generating a single level (or "floor").
type LevelTemplate interface {
	Generate(source *rand.Rand) Map
}

type LevelManager struct {
	filePath string
	levels   map[string]Map

	randomSource *rand.Rand
}

func NewLevelManager(levelFilePath string, seed int64) *LevelManager {
	src := rand.NewSource(seed)
	fmt.Printf("Seed: %d\n", seed)
	return &LevelManager{
		filePath:     levelFilePath,
		levels:       make(map[string]Map),
		randomSource: rand.New(src),
	}
}

func (m *LevelManager) GetLevel(name string) (Map, bool) {
	lvl, ok := m.levels[name]
	return lvl, ok
}

// AddLevel
func (m *LevelManager) AddLevel(name string, template LevelTemplate) Map {
	m.levels[name] = template.Generate(m.randomSource)
	return m.levels[name]
}
