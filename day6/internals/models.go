package internals

type Cell = int

const (
	EMPTY Cell = iota
	OBSTACLE
)

var CellName = map[Cell]string{
	EMPTY:    "empty",
	OBSTACLE: "obstacle",
}

type Coordinate struct {
	x int
	y int
}

type MapMatrix = [][]Cell

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var DirectionName = map[Direction]string{
	UP:    "up",
	DOWN:  "down",
	LEFT:  "left",
	RIGHT: "right",
}

type GuardState struct {
	Coordinate      Coordinate
	FacingDirection Direction
}
