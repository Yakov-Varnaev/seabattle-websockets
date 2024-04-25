package seabattle

import "fmt"

type AlreadyShot struct {
	Cell Cell
}

func (e AlreadyShot) Error() string {
	return fmt.Sprintf("%+v was already shot", e.Cell)
}

type CellIsTaken struct {
	Cell Cell
}

func (e CellIsTaken) Error() string {
	return fmt.Sprintf("%+v is already taken", e.Cell)
}

type NoShipsLeft struct {
	kind  ShipKind
	count int
}

func (e NoShipsLeft) Error() string {
	return fmt.Sprintf("You've used all %d ships of type %s", e.count, e.kind)
}

type BadCells struct {
	Cells []Cell
}

func (e BadCells) Error() string {
	return fmt.Sprintf(
		"Can't place ships on this cells: %v", e.Cells,
	)
}
