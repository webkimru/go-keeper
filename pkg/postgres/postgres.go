package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const _defaultConnTimeout = time.Second * 60

type Postgres struct {
	Pool           *pgxpool.Pool
	ConnectTimeout time.Duration
}

func New(dsn string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		ConnectTimeout: _defaultConnTimeout,
	}
	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	ctx, cancel := context.WithTimeout(context.Background(), pg.ConnectTimeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.New(): %w", err)
	}
	pg.Pool = pool

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
