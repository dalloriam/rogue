package cartography

import (
	"fmt"
	"math/rand"
)

// A LevelTemplate implements logic for generating a single level (or "floor").
type LevelTemplate interface {
	Generate(source *rand.Rand) Map
}

// A LevelManager manages the multiple levels in a game world, and simplifies the transitions between them.
type LevelManager struct {
	filePath string
	levels   map[string]Map

	randomSource *rand.Rand
}

// NewLevelManager initializes & returns a new level manager.
func NewLevelManager(levelFilePath string, seed int64) *LevelManager {
	src := rand.NewSource(seed)
	fmt.Printf("Seed: %d\n", seed)
	return &LevelManager{
		filePath:     levelFilePath,
		levels:       make(map[string]Map),
		randomSource: rand.New(src),
	}
}

// GetLevel returns the desired level given its name.
func (m *LevelManager) GetLevel(name string) (Map, bool) {
	lvl, ok := m.levels[name]
	return lvl, ok
}

// AddLevel adds a new level to the manager.
func (m *LevelManager) AddLevel(name string, template LevelTemplate) Map {
	m.levels[name] = template.Generate(m.randomSource)
	return m.levels[name]
}
