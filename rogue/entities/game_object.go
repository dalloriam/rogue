package entities

import "sync/atomic"

var (
	idCounter uint64
)

// GameObject defines the root game object.
type GameObject struct {
	id uint64

	components map[string]Component

	parent   *GameObject
	children []*GameObject
}

// NewObject returns a new game object.
func NewObject(components ...Component) *GameObject {
	componentMap := make(map[string]Component)
	for _, component := range components {
		componentMap[component.Name()] = component
	}

	return &GameObject{
		id:         atomic.AddUint64(&idCounter, 1),
		components: componentMap,
	}
}

// ID returns the GameObject's ID.
func (o *GameObject) ID() uint64 {
	return o.id
}

// AppendChild appends a child to the current object.
func (o *GameObject) AppendChild(child *GameObject) {
	child.parent = o
	o.children = append(o.children, child)
}

// AddComponents adds components to the entity.
func (o *GameObject) AddComponents(components ...Component) {
	for _, component := range components {
		o.components[component.Name()] = component
	}
}

// HasComponent returns whether the current object has the specified component.
func (o *GameObject) HasComponent(name string) bool {
	_, ok := o.components[name]
	return ok
}
