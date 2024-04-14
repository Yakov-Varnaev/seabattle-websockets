package store

import (
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type LocalStore struct {
	data map[string]*Game
}

func (store *LocalStore) Save(game *Game) (*Game, error) {
	if game.Id == "" {
		id := uuid.NewString()
		game.Id = id
	}

	store.data[game.Id] = game
	slog.Info("Game was saved.", "id", game.Id)
	return game, nil
}

func (store *LocalStore) Get(id string) (*Game, error) {
	game, ok := store.data[id]
	if !ok {
		slog.Error("Game not found.", "id", id)
		return nil, fmt.Errorf("Game not found.")
	}
	slog.Info("Game was succesfully retrieved.", "id", id)
	return game, nil
}
