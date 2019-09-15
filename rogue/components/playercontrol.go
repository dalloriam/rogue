package components

const (
	PlayerControlName = "player_control"
)

type PlayerControl struct{}

func (p *PlayerControl) Name() string {
	return PlayerControlName
}
