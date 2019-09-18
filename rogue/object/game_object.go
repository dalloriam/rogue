package object

import "sync/atomic"

var (
	idCounter uint64
)

type GameObject interface {
	ID() uint64

	AddComponents(c ...Component)
	HasComponent(componentName string) bool
	GetComponent(componentName string) Component
	RemoveComponent(componentName string)
}

// BaseObject defines the root game object.
type BaseObject struct {
	id uint64

	components map[string]Component

	parent   GameObject
	children []GameObject
}

// New returns a new game object.
func New(components ...Component) *BaseObject {
	componentMap := make(map[string]Component)
	for _, component := range components {
		componentMap[component.Name()] = component
	}

	return &BaseObject{
		id:         atomic.AddUint64(&idCounter, 1),
		components: componentMap,
	}
}

// ID returns the BaseObject's ID.
func (o *BaseObject) ID() uint64 {
	return o.id
}

// AddComponents adds components to the entity.
func (o *BaseObject) AddComponents(components ...Component) {
	for _, component := range components {
		o.components[component.Name()] = component
	}
}

func (o *BaseObject) RemoveComponent(name string) {
	delete(o.components, name)
}

// HasComponent returns whether the current object has the specified component.
func (o *BaseObject) HasComponent(name string) bool {
	_, ok := o.components[name]
	return ok
}

func (o *BaseObject) GetComponent(name string) Component {
	// TODO: Make safe
	return o.components[name]
}
