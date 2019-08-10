package connector

import (
	"database/sql"

	_ "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

type CredsKeeper interface {
	ConnectionString() string
}
type Connector struct {
	credsKeeper CredsKeeper
}

func New(ck CredsKeeper) *Connector {
	return &Connector{credsKeeper: ck}
}

func (c *Connector) Connection() (db *sql.DB) {
	db, err := sql.Open("postgres", c.credsKeeper.ConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}
