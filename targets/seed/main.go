package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	. "github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/lib/store/seed"
)

var err error

func main() {
	DB, err = sql.Open("postgres", DBConnectionString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\t\t----- Populating ... -----\n")
	Seed()
	fmt.Printf("\t\t----- Populated! -----\n")
}
