package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"net/http"
	"io"
	"encoding/json"
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
	"map": {
		name:				"map",
		description:	"Displays 20 entries of map-areas - if already used displays the next 20 entries",
		callback: 		locationAreas,
	},
	"mapb": {
		name:				"mapb",
		description: 	"MapBack - Displayed the previous 20 entries",
		callback: 		prevLocationAreas,
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

type LocationAreasResponse struct {
	Count			int					`json:"count"`
	Next			string				`json:"next"`
	Previous 	*string				`json:"previous"`
	Results		[]LocationResult	`json:"results"`
}

type LocationResult struct {
	Name			string				`json:"name"`
	URL			string				`json:"url"`
}


var page int = 0
const pageSize = 20

func locationAreas() {
	page++
	fetchLocationAreas((page - 1) * pageSize)
	
}

func prevLocationAreas() {
	if page > 1 {
		page--
		fetchLocationAreas((page - 1) * pageSize)
	} else {
		fmt.Println("You're on the first page.")
	}
}


func fetchLocationAreas(offset int) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", offset)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	var data LocationAreasResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	for _, loc := range data.Results {
		fmt.Printf("%s\n", loc.Name)
	}
}

