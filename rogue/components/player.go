package components

const (
	PlayerName = "player"
)

type Player struct{}

func (p *Player) Name() string {
	return PlayerName
}
