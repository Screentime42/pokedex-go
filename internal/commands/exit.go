package commands

import (
	"fmt"
	"os"
) 

func commandExit(args []string) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}