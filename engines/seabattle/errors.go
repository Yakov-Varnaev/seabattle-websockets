package seabattle

import "fmt"

type AlreadyShot struct {
	Cell Cell
}

func (e AlreadyShot) Error() string {
	return fmt.Sprintf("%+v was already shot", e.Cell)
}
