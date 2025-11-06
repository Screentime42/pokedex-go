package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
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
		if cmd, ok := commands[cmdName]; ok {
			cmd.callback()
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


type cliCommand struct {
		name			string
		description	string
		callback		func()
}


var commands = map[string]cliCommand {
	"exit": {
		name:				"exit",
		description:	"Exit the Pokedex",
		callback:		commandExit,
	},
	"help": {
		name:				"help",
		description: 	"Displays help commands",
		callback:		help,
	},
}

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func help() {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
}