package components

// Name of the component.
const (
	InitiativeName = "initiative"
)

// The Initiative component represents that this object has the right to move at this frame.
type Initiative struct{}

// Name returns the name of the component.
func (i *Initiative) Name() string { return InitiativeName }
