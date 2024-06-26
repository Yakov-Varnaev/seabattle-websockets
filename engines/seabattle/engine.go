package seabattle

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Engine struct {
	Game *Game
}

func (e *Engine) Shot(targetCell Cell) ([]Cell, error) {
	filledCells := []Cell{}
	turn := e.Game.State.Turn()
	logFields := log.Fields{
		"turn":       turn,
		"targetCell": fmt.Sprintf("%+v", targetCell),
	}
	log.WithFields(logFields).Info("Shooting")
	var field *Field
	if turn == "1" {
		field = e.Game.Field2
	} else {
		field = e.Game.Field1
	}
	_, ok := field.Shots[targetCell]
	// if there is shot then we shouldn't process it
	if ok {
		return filledCells, &AlreadyShot{targetCell}
	}
	log.WithFields(logFields).Info("Hit Cell")
	filledCells = append(filledCells, targetCell)
	field.Shots[targetCell] = true // we should do it in the very end probably

	// check if we hit the ship
	shipsCoords := field.Ships.GetCells()
	ship, ok := shipsCoords[targetCell]
	if !ok {
		log.WithFields(logFields).Info("No ship at the target cell")
		// if there is no ship just return the current result
		e.Game.State.NextTurn()
		return filledCells, nil
	}

	logFields["kind"] = ship.kind
	logFields["direction"] = ship.direction
	log.WithFields(logFields).Info("Hit the ship")
	shipCells := ship.GetCells() // maybe it's better to filter coordinates from shipsCoords?

	isShipDead := true
	for _, coord := range shipCells {
		_, isHit := field.Shots[coord]
		if !isHit {
			isShipDead = false
			break
		}
	}
	logFields["isDead"] = isShipDead
	log.WithFields(logFields).Info("Ship status")

	if isShipDead {
		// fill space around the ship
		log.Info("Filling cells around the ship")
		filled := field.FillAround(ship)
		filledCells = append(filledCells, filled...)
	}

	SortCells(filledCells)

	return filledCells, nil
}
