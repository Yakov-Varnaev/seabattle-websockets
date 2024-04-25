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

var SHIP_AMOUNT_BY_KIND map[ShipKind]int = map[ShipKind]int{
	ShipOne:   4,
	ShipTwo:   3,
	ShipThree: 2,
	ShipFour:  1,
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
	cell      Cell
	direction Direction
}

func (s *Ship) CellsTaken() []Cell {
	shipCoord := s.GetCells()
	c1, c2 := shipCoord[0], shipCoord[len(shipCoord)-1]
	res := []Cell{}
	startX, endX := c1.X-1, c2.X+1
	startY, endY := c1.Y-1, c2.Y+1
	if startX < 0 {
		startX = 0
	}
	if endX > 9 {
		endX = 9
	}
	if startY < 0 {
		startY = 0
	}
	if endY > 9 {
		endY = 9
	}
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			res = append(res, Cell{x, y})
		}
	}
	return res
}

// Returns ship's cells. It's guaranteed that
// cell[i].X <= cell[j].X and cell[i].Y <= cell[j].Y
// where i < j.
func (s Ship) GetCells() []Cell {
	length, ok := SHIP_LENGTH_BY_TYPE[s.kind]
	if !ok {
		panic(fmt.Sprintf("Unknown ShipKind=%s", s.kind))
	}
	cells := make([]Cell, length)
	multiplier := 1
	if s.direction == UP || s.direction == LEFT {
		multiplier = -1
	}
	for i := 0; i < length; i++ {
		if s.direction == UP || s.direction == DOWN {
			// add cells vertically
			cells[i] = Cell{
				X: s.cell.X,
				Y: s.cell.Y + (i * multiplier),
			}
		} else {
			// add cells horizontally
			cells[i] = Cell{
				X: s.cell.X + (i * multiplier),
				Y: s.cell.Y,
			}
		}
	}
	SortCells(cells)
	return cells
}

type Ships []Ship

func (s Ships) GetCells() map[Cell]Ship {
	cells := map[Cell]Ship{}
	for _, ship := range s {
		for _, cell := range ship.GetCells() {
			cells[cell] = ship
		}
	}
	return cells
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

// TODO: write tests, add logging
func (f *Field) PlaceShip(ship Ship) error {
	logger := log.WithFields(
		log.Fields{
			"kind":      ship.kind,
			"cell":      fmt.Sprintf("%+v", ship.cell),
			"direction": ship.direction,
		},
	)
	takenCells := f.Ships.GetCells()
	for _, cell := range ship.CellsTaken() {
		existingShip, isTaken := takenCells[cell]
		if isTaken {
			logger.WithFields(
				log.Fields{
					"existingShip.kind":      existingShip.kind,
					"existsingShip.cell":     existingShip.cell,
					"existingShip.direction": existingShip.direction,
				},
			).Errorf("Cannot place ship as it taken")
			return CellIsTaken{cell}
		}
	}
	shipKindCnt := 0
	for _, existingShip := range f.Ships {
		if ship.kind == existingShip.kind {
			shipKindCnt++
		}
	}
	allowedAmount, ok := SHIP_AMOUNT_BY_KIND[ship.kind]
	if !ok {
		panic(fmt.Sprintf("Unknown ship.kind=%s", ship.kind))
	}
	if allowedAmount == shipKindCnt {
		return NoShipsLeft{kind: ship.kind, count: shipKindCnt}
	} else if allowedAmount < shipKindCnt {
		panic("We placed more ships than allowed")
	}
	f.Ships = append(f.Ships, ship)
	logger.Info("Ship is successfully placed")
	return nil
}

// Fill rectangular area from c1.X, c1.Y to c2.X, c2.Y
func (f *Field) FillRect(c1, c2 Cell) []Cell {
	cells := []Cell{}
	if c1.X > c2.X || c1.Y > c2.Y {
		panic("c1 must be less than c2")
	}
	for x := c1.X; x <= c2.X; x++ {
		for y := c1.Y; y <= c2.Y; y++ {
			c := Cell{x, y}
			_, wasFilled := f.Shots[c]
			f.Shots[c] = true
			if !wasFilled {
				cells = append(cells, c)
			}
		}
	}
	return cells
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
