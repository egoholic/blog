package main

import (
	"database/sql"
	"fmt"

	. "github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/lib/store/schema"
)

func main() {
	fmt.Printf("\n\t\t----- Starting migration ... -----\n")
	db, err := sql.Open("postgres", DBConnectionStringWithoutDB)
	defer db.Close()
	if err != nil {
		return
	}
	err = Apply(db)
	if err != nil {
		fmt.Printf("Error occured during migrating: `%s`\n", err.Error())
		panic(err)
	}
	fmt.Printf("\t\t----- Migrating succeed! -----\n")
}
