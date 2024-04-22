package seabattle

import "errors"

var AlreadyShot error = errors.New("Cell is already shot.")

type Engine struct {
	Game *Game
}

func (e *Engine) Shot(targetCell Cell) error {
	turn := e.Game.State.Turn()
	var field *Field
	if turn == "1" {
		field = e.Game.Field2
	} else {
		field = e.Game.Field1
	}

	_, ok := field.Shots[targetCell]
	// if there is shot then we shouldn't process it
	if ok {
		return AlreadyShot
	}
	field.Shots[targetCell] = true // we should do it in the very end probably

	// check if we hit the ship

	shipsCoords := field.Ships.GetCoordinates()

	ship, ok := shipsCoords[targetCell]
	if !ok {
		// if there is no ship just return the current result
		return nil
	}

	shipCoords := ship.GetCoordinates() // maybe it's better to filter coordinates from shipsCoords?

	isShipDead := true
	for _, coord := range shipCoords {
		_, isHit := field.Shots[coord]
		if !isHit {
			isShipDead = false
			break
		}
	}

	if isShipDead {
		// fill space around the ship
	}

	return nil
}
