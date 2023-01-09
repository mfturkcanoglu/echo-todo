package client

import "database/sql"

type Client interface {
	ConnectDb() error
	RunMigrations()
	Db() *sql.DB
	Close()
}
