package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Configuration struct {
	dbCredentials *DBCredentials
}

type DBCredentials struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var Config *Configuration

func init() {
	Config = &Configuration{&DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: ""}}
}

func (c *DBCredentials) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s pasword=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBName)
}

func (c *Configuration) DBCredentials() *DBCredentials {
	return c.dbCredentials
}

func (c *Configuration) DB() *sql.DB {
	db, err := sql.Open("postgres", Config.DBCredentials().ConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}
