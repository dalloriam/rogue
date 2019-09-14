package cartography

// A LevelTemplate implements logic for generating a single level (or "floor").
type LevelTemplate interface {
	Generate() Map
}

type LevelManager struct {
	filePath string
	levels   map[string]Map
}

func NewLevelManager(levelFilePath string) *LevelManager {
	return &LevelManager{
		filePath: levelFilePath,
		levels:   make(map[string]Map),
	}
}

func (m *LevelManager) Save() error {
	return nil
}

func (m *LevelManager) Load() error {
	return nil
}

// AddLevel
func (m *LevelManager) AddLevel(name string, template LevelTemplate) {
	m.levels[name] = template.Generate()
}
