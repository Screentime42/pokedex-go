package commands

type CliCommand struct {
		Name			string
		Description	string
		Callback		func(args []string)
}


var CommandMap = map[string]CliCommand {
	"exit": {
		Name:				"exit",
		Description:	"Exit the Pokedex",
		Callback:		commandExit,
	},
	"help": {
		Name:				"help",
		Description: 	"Displays help commands",
		Callback:		help,
	},
	"map": {
		Name:				"map",
		Description:	"Displays 20 entries of map-areas - if already used displays the next 20 entries",
		Callback: 		locationAreas,
	},
	"mapb": {
		Name:				"mapb",
		Description: 	"MapBack - Displayed the previous 20 entries",
		Callback: 		prevLocationAreas,
	},
	"explore": {
		Name:				"explore",
		Description:	"Shows all the Pokemon available to be encountered within the area",
		Callback:		commandExplore,
	},
	"catch": {
		Name:				"catch",
		Description:	"Attempt to catch a pokemon.",
		Callback:		commandCatch,
	},
}