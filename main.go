package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"

	cmds "pokedex-go/internal/commands"
)

func main() {
   scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to your Pokedex. Type exit to quit.")

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "Incorrect input.")
			break
		}

		input := strings.TrimSpace(scanner.Text())
		words := cleanInput(input)
		cmdName := words[0]

		if cmd, ok := cmds.CommandMap[cmdName]; ok {
			cmd.Callback()
		} else {
			fmt.Println("Unknown command")
		}

		

		if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Error reading input:", err)
    }
	}
}


func cleanInput(text string) []string {
	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return words
}








