package logga

import (
	"os"

	"github.com/rs/zerolog"
)

type Logga struct {
	Lg zerolog.Logger
}

func New() *Logga {

	logger := zerolog.New(os.Stderr)

	logga := Logga{
		Lg: logger,
	}

	return &logga
}

func (l *Logga) Info(msg string) {
	l.Lg.Info().Msg(msg)
}

func (l *Logga) Error(msg string) {
	l.Lg.Error().Msg(msg)
}
