package env

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"

	"cryptowatcher.example/internal/pkg/logga"
)

type Env struct {
	l *logga.Logga
}

func New(l *logga.Logga) (*Env, error){

	err := LoadEnvironmentVars()
	if err != nil {
		return nil, err
	}

	e := Env{
		l: l,
	}

	return &e, nil
}

func (e *Env) GetVar(v string) string {

	return os.Getenv(v)
}

func (e *Env) TestEnvironmentVars(requiredVars []string) error {

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("Variable not found in environment file: %s\n", v)
		}
	}

	return nil
}

func LoadEnvironmentVars() error {

	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)

	var envFile string // We need to determine absolute path for testing.

	if os.Getenv("APP_ENV") == "test" {
		envFile = basepath +"/../../.env.testing"
	} else {
		envFile = basepath +"/../../.env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("Could not load environment file: %s\n", envFile)
	}

	return nil
}

