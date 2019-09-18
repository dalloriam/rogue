package components

const (
	FocusName = "focus"
)

type Focus struct {
	// Punctual indicates that the camera should focus on this object only once.
	Punctual bool

	// A higher priority means that the camera should prioritize focusing on this entity.
	Priority int
}

// Name returns this component's name.
func (f *Focus) Name() string {
	return FocusName
}
