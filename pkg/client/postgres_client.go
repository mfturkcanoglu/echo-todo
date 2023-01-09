package client

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "todo"
)

type postgresClient struct {
	log *zap.Logger
	db  *sql.DB
}

func NewPostgresClient(l *zap.Logger) Client {
	return &postgresClient{
		log: l,
	}
}

func (c *postgresClient) ConnectDb() (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	c.log.Log(zap.DebugLevel, "Successfully connected!")
	// Assign db to object
	c.db = db
	c.RunMigrations()
	return
}

func (c *postgresClient) RunMigrations() {
	c.log.Log(zap.DebugLevel, "No migration found")
}

func (c *postgresClient) Close() {
	c.db.Close()
}

func (c *postgresClient) Db() *sql.DB {
	err := c.ConnectDb()
	if err != nil {
		c.log.Fatal("Cannot connect to db")
		return nil
	}
	return c.db
}
