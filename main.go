package main

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"os"
	"strings"
	"terminal_ai/ai_models"
	"terminal_ai/settings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(1)
	}

	prompt := strings.Join(args, " ")
	settings := settings.ReadSettings("settings.json")

	ai, error := ai_models.NewGemini(settings.Key)
	if error != nil {
		panic(error)
	}

	waitMessage, error := glamour.Render("**Pensando...**", settings.Theme)
	if error != nil {
		panic(error)
	}

	fmt.Println(waitMessage)

	response, error := ai.GetResponse(prompt, settings.MaxTokens)
	if error != nil {
		panic(error)
	}

	response, error = glamour.Render(response, settings.Theme)
	if error != nil {
		panic(error)
	}

	fmt.Println(response)
}
