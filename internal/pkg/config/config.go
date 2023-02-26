package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
)

type Config struct {
	Aggregator Aggregator `json:"aggregator"`
	DB         DB         `json:"database"`
}

type Aggregator struct {
	Name string `json:"name"`
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
		l.Error().Msgf("unable to read config file. %s", err.Error())

		return nil, fmt.Errorf("unable to read config file `%s`. %w", configFile, err)
	}

	if err = json.Unmarshal(f, &c); err != nil {
		return nil, err
	}

	if c.DB.Host == "" || c.DB.Name == "" || c.DB.User == "" || c.DB.Pass == "" {
		return &c, fmt.Errorf("some configuration is missing")
	}

	return &c, nil
}
