package seabattle

import "fmt"

type Direction string

type SeabattleMap [10][10]int

const (
	UP    Direction = "up"
	DOWN            = "down"
	RIGHT           = "right"
	LEFT            = "left"
)

type ShipKind string

// TODO: better names?
const (
	ShipOne   ShipKind = "one"
	ShipTwo            = "two"
	ShipThree          = "three"
	ShipFour           = "four"
)

var SHIP_LENGTH_BY_TYPE map[ShipKind]int = map[ShipKind]int{
	ShipOne:   1,
	ShipTwo:   2,
	ShipThree: 3,
	ShipFour:  4,
}

type Cell struct {
	X int
	Y int
}

type Ship struct {
	kind      ShipKind
	coord     Cell
	direction Direction
}

// Returns ship's coordinates
func (s Ship) GetCoordinates() []Cell {
	length, ok := SHIP_LENGTH_BY_TYPE[s.kind]
	if !ok {
		panic(fmt.Sprintf("Unknown ShipKind=%s", s.kind))
	}
	coordinates := make([]Cell, length)
	multiplier := 1
	if s.direction == UP || s.direction == LEFT {
		multiplier = -1
	}
	for i := 0; i < length; i++ {
		if s.direction == UP || s.direction == DOWN {
			// add cells vertically
			coordinates[i] = Cell{
				X: s.coord.X,
				Y: s.coord.Y + (i * multiplier),
			}
		} else {
			// add cells horizontally
			coordinates[i] = Cell{
				X: s.coord.X + (i * multiplier),
				Y: s.coord.Y,
			}
		}
	}
	return coordinates
}

type Ships []*Ship

func (s Ships) GetCoordinates() map[Cell]*Ship {
	coordinates := map[Cell]*Ship{}
	for _, ship := range s {
		for _, coord := range ship.GetCoordinates() {
			coordinates[coord] = ship
		}
	}
	return coordinates
}

type Field struct {
	Shots map[*Cell]bool
	Ships Ships
}

// This method will return the matrix which represents
// the game field as a game map.
func (f *Field) AsMap() SeabattleMap {
	m := SeabattleMap{}
	//TODO: Finish this method
	return m
}

type State struct {
	turn string
}

// Sets the current active player.
func (s *State) SetTurn(player string) {
	s.turn = player
}

// Returns the id of the active player.
func (s *State) Turn() string {
	return s.turn
}

type Game struct {
	State  *State
	Id     string
	Field1 *Field
	Field2 *Field
}
