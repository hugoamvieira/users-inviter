package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config is the structure that holds this program's configuration.
// You just need to specify it in a JSON file, pass it onto NewJSONConfig and it'll read it in.
type Config struct {
	Lat             float64 `json:"latitude"`
	Lng             float64 `json:"longitude"`
	DistThresholdKm float64 `json:"distance_threshold_km"`
	UsersFilePath   string  `json:"users_file_path"`
}

// NewJSON receives the filepath for the configuration JSON file, unmarshals it and
// returns the internal structure that represents the configuration.
func NewJSON(filePath string) (*Config, error) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
