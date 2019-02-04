package postgresql

import (
	"context"
	"database/sql"

	"github.com/lib/pq"

	"github.com/moonrhythm/archer/core"
)

var _ = pq.Driver{}

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
	_, err := s.db.ExecContext(ctx, `
		insert into storage
			(key, value)
		values
			($1, $2)
		on conflict do update set
			value = excluded.value
	`, key, value)
	return err
}

func (s *storage) Get(ctx context.Context, key string) (value []byte, err error) {
	err = s.db.QueryRowContext(ctx, `
		select value
		from storage
		where key = $1
	`, key).Scan(&value)
	if err == sql.ErrNoRows {
		err = core.ErrNotFound
	}
	return
}

func (s *storage) Del(ctx context.Context, key string) error {
	_, err := s.db.ExecContext(ctx, `
		delete from storage
		where key = $1
	`, key)
	return err
}

func (s *storage) Keys(ctx context.Context, prefix string) (keys []string, err error) {
	err = s.db.QueryRowContext(ctx, `
		select array_agg(key)
		from storage
		where key like $1
	`, prefix+"%").Scan(pq.Array(&keys))
	return
}

func (s *storage) Exists(ctx context.Context, key string) (exists bool, err error) {
	err = s.db.QueryRowContext(ctx, `
		select exists (
		    select 1
		    from storage
		    where key = $1
		)
	`, key).Scan(&exists)
	return
}
