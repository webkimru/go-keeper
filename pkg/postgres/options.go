package postgres

import "time"

type Option func(*Postgres)

func ConnectTimeout(timeout int) Option {
	return func(p *Postgres) {
		p.ConnectTimeout = time.Duration(timeout) * time.Second
	}
}
