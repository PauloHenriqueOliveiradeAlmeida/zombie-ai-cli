package settings

import (
	"encoding/json"
	"os"
)

type Settings struct {
	Key       string `json:"key"`
	MaxTokens int    `json:"max_tokens"`
	Theme     string `json:"theme"`
}

func ReadSettings(location string) Settings {
	file, error := os.Open(location)
	if error != nil {
		panic(error)
	}

	settings := Settings{}

	decoder := json.NewDecoder(file)
	error = decoder.Decode(&settings)
	if error != nil {
		panic(error)
	}

	return settings
}
