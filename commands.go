package main

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	Next     any
	Previous any
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
			description: "Lists 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists previous 20 areas",
			callback:    commandMapb,
		},
	}

}
