package settings

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type Settings struct {
	Key       string `json:"key"`
	MaxTokens int    `json:"max_tokens"`
	Theme     string `json:"theme"`
}

func (this *Settings) Write(location string, filename string) error {
	error := os.MkdirAll(location, 0700)
	if error != nil {
		return error
	}

	path := filepath.Join(location, filename)

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return json.NewEncoder(file).Encode(this)
}

func ReadSettings(location string, filename string) (*Settings, error) {
	file, error := os.Open(filepath.Join(location, filename))
	if error != nil {
		return nil, errors.New("Configurações não encontradas, use 'zombie configure' para configurar")
	}

	settings := Settings{}

	decoder := json.NewDecoder(file)
	error = decoder.Decode(&settings)
	if error != nil {
		panic(error)
	}

	return &settings, nil
}

func GetSettingsPath() (string, string, error) {
	home, error := os.UserHomeDir()
	if error != nil {
		return "", "", error
	}

	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("AppData"), "zombie"), "settings.json", nil
	default:
		return filepath.Join(home, ".config", "zombie"), "settings.json", nil
	}
}
