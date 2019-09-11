package main

import (
	"fmt"

	. "github.com/egoholic/blog/lib/store/schema"
)

func main() {
	fmt.Printf("\n\t\t----- Starting migration ... -----\n")
	err := Apply()
	if err != nil {
		fmt.Printf("Error occured during migrating: `%s`\n", err.Error())
		panic(err)
	}
	fmt.Printf("\t\t----- Migrating succeed! -----\n")
}
