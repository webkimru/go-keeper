package postgres

import "time"

// Option is functional options for PostgreSQL.
type Option func(*Postgres)

// ConnectTimeout is the time limit for connection timeout,
func ConnectTimeout(timeout int) Option {
	return func(p *Postgres) {
		p.ConnectTimeout = time.Duration(timeout) * time.Second
	}
}
