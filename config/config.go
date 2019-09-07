package config

import (
	"errors"
	"fmt"
	"strings"

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
	Config = &Configuration{&DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: "blog"}}
}

func (c *DBCredentials) ConnectionString() (string, error) {
	var sb strings.Builder
	if len(c.Host) > 0 {
		sb.WriteString(fmt.Sprintf("host=%s ", c.Host))
	}
	sb.WriteString(fmt.Sprintf("port=%d ", c.Port))
	if len(c.Password) > 0 {
		sb.WriteString(fmt.Sprintf("password=%s ", c.Password))
	}
	if len(c.User) > 0 {
		sb.WriteString(fmt.Sprintf("user=%s ", c.User))
	}
	if len(c.DBName) > 0 {
		sb.WriteString(fmt.Sprintf("dbname=%s ", c.DBName))
	} else {
		return "", errors.New("no db name given")
	}
	sb.WriteString("sslmode=disable")
	return sb.String(), nil
}

func (c *DBCredentials) ConnectionStringWithoutDB() (string, error) {
	var sb strings.Builder
	if len(c.Host) > 0 {
		sb.WriteString(fmt.Sprintf("host=%s ", c.Host))
	}
	sb.WriteString(fmt.Sprintf("port=%d ", c.Port))
	if len(c.Password) > 0 {
		sb.WriteString(fmt.Sprintf("password=%s ", c.Password))
	}
	if len(c.User) > 0 {
		sb.WriteString(fmt.Sprintf("user=%s ", c.User))
	}
	sb.WriteString("sslmode=disable")
	return sb.String(), nil
}
func (c *Configuration) DBCredentials() *DBCredentials {
	return c.dbCredentials
}
