package components

// Name of the component.
const (
	PlayerName = "player"
)

// The Player component indicates that the object is player-controlled.
type Player struct{}

// Name returns the component's name.
func (p *Player) Name() string {
	return PlayerName
}
