package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

const _defaultConnTimeout = time.Second

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(dsn string) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _defaultConnTimeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres - New - pgxpool.New(): %w", err)
	}

	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
