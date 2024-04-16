package store

import (
	"fmt"
	"log/slog"
	"time"
)

type Cell struct {
	ShipID int
	IsShot bool
}

type Field struct {
	Data [10][10]*Cell
}

func NewField() *Field {
	field := Field{}
	for i := range 10 {
		for j := range 10 {
			field.Data[i][j] = &Cell{ShipID: 0, IsShot: false}
		}
	}
	slog.Info("Field successfully generated.")
	return &field
}

type Game struct {
	Id      string
	Field1  *Field
	Field2  *Field
	Created time.Time
}

func (game *Game) GetPlayerLinks() map[string]string {
	return map[string]string{
		"player1": fmt.Sprintf("/game/%s/player/1", game.Id),
		"player2": fmt.Sprintf("/game/%s/player/2", game.Id),
	}
}

func NewGame() *Game {
	slog.Info("Creating new game.")
	game := Game{
		Field1:  NewField(),
		Field2:  NewField(),
		Created: time.Now(),
	}
	slog.Info("Game succuessfully created.")
	return &game
}
