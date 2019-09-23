package components

const (
	InitiativeName = "initiative"
)

type Initiative struct{}

func (i *Initiative) Name() string { return InitiativeName }
