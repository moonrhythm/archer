package postgresql

import (
	"context"
	"database/sql"

	"github.com/moonrhythm/archer/core"
)

// New creates new storage implements using postgres
func New(db *sql.DB) core.Storage {
	db.Exec(`
		create table if not exists storage (
			key varchar,
			value varchar,
			primary key (key)
		)
	`)
	return &storage{db}
}

type storage struct {
	db *sql.DB
}

func (s *storage) Set(ctx context.Context, key string, value []byte) error {
	panic("implement me")
}

func (s *storage) Get(ctx context.Context, key string) ([]byte, error) {
	panic("implement me")
}

func (s *storage) Del(ctx context.Context, key string) error {
	panic("implement me")
}

func (s *storage) Keys(ctx context.Context, prefix string) ([]string, error) {
	panic("implement me")
}

func (s *storage) Exists(ctx context.Context, key string) (bool, error) {
	panic("implement me")
}
