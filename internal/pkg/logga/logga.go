package logga

import (
	"os"

	"github.com/rs/zerolog"
)

// todo - replace this with the commented out code below and the new slog Go library.

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

//package logga
//
//import (
//"log"
//"os"
//)
//
//type Logga struct {
//	infoLogger  *log.Logger
//	warnLogger  *log.Logger
//	errorLogger *log.Logger
//}
//
//func New() *Logga {
//
//	flags := log.LstdFlags | log.Lshortfile
//	infoLogger := log.New(os.Stdout, "INFO", flags)
//	warnLogger := log.New(os.Stdout, "WARN", flags)
//	errorLogger := log.New(os.Stdout, "ERROR", flags)
//
//	return &Logga{
//		infoLogger:  infoLogger,
//		warnLogger:  warnLogger,
//		errorLogger: errorLogger,
//	}
//}
//
//func (l *Logga) Info(msg string) {
//	l.infoLogger.Println(msg)
//}
//
//func (l *Logga) Warn(msg string) {
//	l.warnLogger.Println(msg)
//}
//
//func (l *Logga) Error(msg string) {
//	l.errorLogger.Println(msg)
//}
