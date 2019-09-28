package config

import (
	"fmt"
	"strings"

	"github.com/egoholic/cfg"
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

var (
	DBConnectionString          string
	DBConnectionStringWithoutDB string

	Port   int
	dbHost string
	dbPort int
	dbUser string
	dbPwd  string
	DBName string
	config *cfg.Cfg
	err    error
)

func init() {
	defaults := map[string]interface{}{}
	defaults["dbhost"] = "localhost"
	defaults["dbport"] = "5432"
	defaults["dbuser"] = "postgres"
	defaults["dbname"] = "stoa_blogging_development"
	defaults["dbpwd"] = ""
	defaults["port"] = "3000"
	config := cfg.Config(defaults)
	Port, err = config.IntArg("Web server port", "The port which web server listens to, like: 3000.", "port")
	if err != nil {
		panic(err)
	}
	dbHost, err = config.StringArg("DB Host name", "Database connection host name, like: 'localhost'.", "dbhost")
	if err != nil {
		panic(err)
	}
	dbPort, err = config.IntArg("DB port", "Database connection port, like: 5432.", "dbport")
	if err != nil {
		panic(err)
	}
	dbUser, err = config.StringArg("DB User", "Database connection user name, like: 'postgres'.", "dbuser")
	if err != nil {
		panic(err)
	}
	dbPwd, err = config.StringArg("DB Password", "Database connection password.", "dbpwd")
	if err != nil {
		panic(err)
	}
	DBName, err = config.StringArg("Database name", "Database name, like: 'stoa_blogging_development'.", "dbname")
	if err != nil {
		panic(err)
	}
	config.AddHelpCommand()
	DBConnectionString = genConnectionString(true)
	DBConnectionStringWithoutDB = genConnectionString(false)
}

func genConnectionString(withDBName bool) string {
	var sb strings.Builder
	if len(dbHost) > 0 {
		sb.WriteString(fmt.Sprintf("host=%s ", dbHost))
	}
	sb.WriteString(fmt.Sprintf("port=%d ", dbPort))
	if len(dbPwd) > 0 {
		sb.WriteString(fmt.Sprintf("password=%s ", dbPwd))
	}
	if len(dbUser) > 0 {
		sb.WriteString(fmt.Sprintf("user=%s ", dbUser))
	}
	if withDBName {
		if len(DBName) > 0 {
			sb.WriteString(fmt.Sprintf("dbname=%s ", DBName))
		}
	}
	sb.WriteString("sslmode=disable")
	return sb.String()
}
