package main

import (
	"context"
	"fmt"

	. "github.com/egoholic/blog/store/schema"
)

func main() {
	fmt.Printf("\n\t\t----- Starting migration ... -----\n")
	ctx := context.Background()
	err := Apply(ctx)
	if err != nil {
		fmt.Printf("Error occured during migrating: `%s`\n", err.Error())
		panic(err)
	}
	fmt.Printf("\t\t----- Migrating succeed! -----\n")
}
