package cartography

import "math/rand"

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
	return &LevelManager{
		filePath:     levelFilePath,
		levels:       make(map[string]Map),
		randomSource: rand.New(src),
	}
}

func (m *LevelManager) GetLevel(name string) Map {
	return m.levels[name]
}

func (m *LevelManager) Save() error {
	return nil
}

func (m *LevelManager) Load() error {
	return nil
}

// AddLevel
func (m *LevelManager) AddLevel(name string, template LevelTemplate) {
	m.levels[name] = template.Generate(m.randomSource)
}
