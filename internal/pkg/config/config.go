package config

import (
	"cryptowatcher.example/internal/pkg/logga"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Tracker Tracker `json:"tracker"`
	DB      DB      `json:"database"`
}

type Tracker struct {
	APIKey string `json:"APIKey"`
	Host   string `json:"host"`
}

type DB struct {
	User string `json:"user"`
	Pass string `json:"password"`
	Host string `json:"host"`
	Name string `json:"name"`
}

func New(lg *logga.Logga, configFile string) (*Config, error) {
	l := lg.Lg.With().Str("config", "New").Logger()
	l.Info().Msgf("Loading config from file: %s", configFile)

	var c Config

	f, err := os.ReadFile(configFile)
	if err != nil {
		l.Error().Msgf("unable to read config file. %w", err.Error())
		return nil, fmt.Errorf("unable to read config file `%s`. %w", configFile, err)
	}

	err = json.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	if c.DB.Host == "" || c.DB.Name == "" || c.DB.User == "" || c.DB.Pass == "" || c.Tracker.Host == "" || c.Tracker.APIKey == "" {
		return &c, fmt.Errorf("some configuration is missing")
	}

	return &c, nil
}
