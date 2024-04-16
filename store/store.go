package store

import "log/slog"

type Store interface {
	Save(game *Game) (*Game, error)
	Get(id string) (*Game, error)
	List() ([]*Game, error)
}

var store Store

func Init(storeType string) {
	slog.Info("Initializng store", "type", storeType)
	store = LocalStore{
		data: map[string]*Game{},
	}
	slog.Info("Store initialized")
}

func GetStore() Store {
	if store == nil {
		panic("Store was not initialized yet.")
	}
	return store
}
