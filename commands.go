package main

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, arg string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Lists the available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous page of location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Lists all the pokemon located in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempts to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Shows the details of a caught pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all caught pokemon",
			callback:    commandPokedex,
		},
	}

}
