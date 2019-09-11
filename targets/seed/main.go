package main

import (
	"fmt"

	. "github.com/egoholic/blog/lib/store/seed"
)

func main() {
	fmt.Printf("\n\t\t----- Starting populating ... -----\n")
	err := Seed()
	if err != nil {
		fmt.Printf("Error occured during DB populating: `%s`\n", err.Error())
		panic(err)
	}
	fmt.Printf("\t\t----- Populating succeed! -----\n")
}
