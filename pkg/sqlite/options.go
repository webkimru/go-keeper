package sqlite

import "time"

// Option is functional options for SQLite.
type Option func(*SQLite)

// PingInterval is the time interval for ping SQLite.
func PingInterval(interval int) Option {
	return func(p *SQLite) {
		p.PingInterval = time.Duration(interval) * time.Second
	}
}

// DataSourcePath is the database source path.
func DataSourcePath(path string) Option {
	return func(p *SQLite) {
		p.DataSourcePath = path
	}
}
