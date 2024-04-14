package store

type Store interface {
	Save(game *Game) (*Game, error)
	Get(id string) (*Game, error)
}
