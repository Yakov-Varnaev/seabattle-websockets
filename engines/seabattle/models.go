package seabattle

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sort"
)

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

func SortCells(a []Cell) {
	sort.Slice(a, func(i, j int) bool {
		ci, cj := a[i], a[j]
		if ci.X == cj.X {
			return ci.Y < cj.Y
		}
		return ci.X < cj.X
	})
}

type Ship struct {
	kind      ShipKind
	coord     Cell
	direction Direction
}

// Returns ship's coordinates. It's guaranteed that
// cell[i].X <= cell[j].X and cell[i].Y <= cell[j].Y
// where i < j.
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
	SortCells(coordinates)
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
	Shots map[Cell]bool
	Ships Ships
}

func NewField() *Field {
	return &Field{
		Shots: map[Cell]bool{},
		Ships: Ships{},
	}
}

func (f *Field) FillRect(c1, c2 Cell) []Cell {
	coordinates := []Cell{}
	if c1.X > c2.X || c1.Y > c2.Y {
		panic("c1 must be less than c2")
	}
	for x := c1.X; x <= c2.X; x++ {
		for y := c1.Y; y <= c2.Y; y++ {
			c := Cell{x, y}
			_, wasFilled := f.Shots[c]
			f.Shots[c] = true
			if !wasFilled {
				coordinates = append(coordinates, c)
			}
		}
	}
	return coordinates
}

// Check if the field is valid.
// The field is considered to be valid if:
//   - There is at least one cell between each ship
//   - There is exactly(ship - amount): 4 - 1, 3 - 2, 2 - 3, 1 - 4
func (f *Field) Validate() error {
	return nil
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

func (s *State) NextTurn() {
	log.Info("Switching turn")
	if s.turn == "1" {
		s.turn = "2"
	} else {
		s.turn = "1"
	}
}

type Game struct {
	State  *State
	Id     string
	Field1 *Field
	Field2 *Field
}

func NewGame() *Game {
	return &Game{
		State: &State{
			turn: "1",
		},
		Id:     "1",
		Field1: NewField(),
		Field2: NewField(),
	}
}
