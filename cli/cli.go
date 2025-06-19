package cli

import (
	"errors"
	"fmt"
	"os"
	"terminal_ai/ai_models"
	"terminal_ai/settings"
	"time"

	"github.com/charmbracelet/glamour"
)

type Color string

const (
	normalBlack  Color = "\033[0;30m"
	normalRed    Color = "\033[0;31m"
	normalGreen  Color = "\033[0;32m"
	normalYellow Color = "\033[0;33m"
	normalBlue   Color = "\033[0;34m"
	normalPurple Color = "\033[0;35m"
	normalCyan   Color = "\033[0;36m"
	normalWhite  Color = "\033[0;37m"
)

const (
	boldBlack  Color = "\033[1;30m"
	boldRed    Color = "\033[1;31m"
	boldGreen  Color = "\033[1;32m"
	boldYellow Color = "\033[1;33m"
	boldBlue   Color = "\033[1;34m"
	boldPurple Color = "\033[1;35m"
	boldCyan   Color = "\033[1;36m"
	boldWhite  Color = "\033[1;37m"
)

const reset Color = "\033[0m"

const asciiArt string = `
	vamos começar!
  ▒███████▒ ▒█████   ███▄ ▄███▓ ▄▄▄▄    ██▓▓█████ 
  ▒ ▒ ▒ ▄▀░▒██▒  ██▒▓██▒▀█▀ ██▒▓█████▄ ▓██▒▓█   ▀ 
  ░ ▒ ▄▀▒░ ▒██░  ██▒▓██    ▓██░▒██▒ ▄██▒██▒▒███   
    ▄▀▒   ░▒██   ██░▒██    ▒██ ▒██░█▀  ░██░▒▓█  ▄ 
  ▒███████▒░ ████▓▒░▒██▒   ░██▒░▓█  ▀█▓░██░░▒████▒
  ░▒▒ ▓░▒░▒░ ▒░▒░▒░ ░ ▒░   ░  ░░▒▓███▀▒░▓  ░░ ▒░ ░
  ░░▒ ▒ ░ ▒  ░ ▒ ▒░ ░  ░      ░▒░▒   ░  ▒ ░ ░ ░  ░
  ░ ░ ░ ░ ░░ ░ ░ ▒  ░      ░    ░    ░  ▒ ░   ░   
    ░ ░        ░ ░         ░    ░       ░     ░  ░
  ░                                  ░            
`

func render(message string, color Color) string {
	return string(color) + message + string(reset)
}

func Configure() error {
	fmt.Println(render(asciiArt, boldCyan))

	var apiKey string
	fmt.Println(render("Digite sua chave de API: ", boldPurple))
	fmt.Scanln(&apiKey)

	var maxTokens int
	fmt.Println(render("Digite o número máximo de tokens: ", boldPurple))
	fmt.Scanln(&maxTokens)

	var theme string
	fmt.Println(render("Digite o tema: ", boldPurple))
	fmt.Println(render(`
		[ 1 ] - Light
		[ 2 ] - Dark - Padrão
		[ 3 ] - Nightly
		`, boldCyan))

	fmt.Scanln(&theme)
	switch theme {
	case "1":
		theme = "light"
	case "2":
		theme = "dark"
	case "3":
		theme = "nightly"
	default:
		theme = "dark"
	}

	settingsToSave := settings.Settings{
		Key:       apiKey,
		MaxTokens: maxTokens,
		Theme:     theme,
	}

	settingsPath, settingsFilename, error := settings.GetSettingsPath()
	if error != nil {
		return errors.New(render(error.Error(), boldRed))
	}

	settingsToSave.Write(settingsPath, settingsFilename)
	fmt.Println(render("Configurações salvas com sucesso!", boldGreen))
	fmt.Println(render("Use o comando 'zombie' para começar a conversar!", boldGreen))
	os.Exit(0)

	return nil
}

func Ask(prompt string) (string, error) {
	settingsPath, settingsFilename, error := settings.GetSettingsPath()
	if error != nil {
		return "", errors.New(render(error.Error(), boldRed))
	}

	settings, error := settings.ReadSettings(settingsPath, settingsFilename)
	if error != nil {
		return "", errors.New(render(error.Error(), boldRed))
	}

	ai, error := ai_models.NewGemini(settings.Key)
	if error != nil {
		return "", errors.New(render(error.Error(), boldRed))
	}

	done := make(chan struct{})
	go waitAnimation(done)

	response, error := ai.GetResponse(prompt, settings.MaxTokens)
	if error != nil {
		close(done)
		return "", errors.New(render(error.Error(), boldRed))
	}

	response, error = glamour.Render(response, settings.Theme)
	if error != nil {
		close(done)
		return "", errors.New(render(error.Error(), boldRed))
	}

	close(done)
	return response, nil
}

func waitAnimation(done <-chan struct{}) {
	waitingMessages := [4]string{
		"| Estou pensando.",
		"/ Estou pensando..",
		"- Estou pensando...",
		"\\ Estou pensando...",
	}

	index := 0
	for {
		select {
		case <-done:
			fmt.Print("\r                                  \r")
			return
		default:
			fmt.Printf("\r%s", render(waitingMessages[index%len(waitingMessages)], boldWhite))
			index++
			time.Sleep(200 * time.Millisecond)
		}
	}
}
