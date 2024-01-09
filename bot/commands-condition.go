package bot

import (
	"strings"
)

type commandCondition struct {
	userCommand string
}

func (command commandCondition) isGreetCommand() (bool) {
	greetCommands := []string {
		"hi bot",
		"hi! bot",
		"hello bot",
		"hello! bot",
	}

	for _, greetCommand := range greetCommands {
		if greetCommand == command.userCommand{
			return true
		}
	}
	return false;
}

func (command commandCondition) isCityWeatherCommand() (bool) {
	return strings.Contains(command.userCommand, "!city")
}
