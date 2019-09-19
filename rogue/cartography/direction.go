package cartography

type Direction string

const (
	DirectionUp        Direction = "up"
	DirectionDown      Direction = "down"
	DirectionLeft      Direction = "left"
	DirectionRight     Direction = "right"
	DirectionDownRight Direction = "downright"
	DirectionDownLeft  Direction = "downleft"
	DirectionUpLeft    Direction = "upleft"
	DirectionUpRight   Direction = "upright"
	NoDirection        Direction = "none"
)
