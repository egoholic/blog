package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	. "github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/lib/store/seed"
)

func main() {
	connStr, err := Config.DBCredentials().ConnectionString()
	if err != nil {
		panic(err)
	}
	DB, err = sql.Open("postgres", connStr)
	//	defer DB.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\t\t----- Starting populating ... -----\n")
	Seed()
	if err != nil {
		fmt.Printf("Error occured during DB populating: `%s`\n", err.Error())
		panic(err)
	}
	fmt.Printf("\t\t----- Populating succeed! -----\n")
}
