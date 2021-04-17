package env

import (
	"cryptowatcher.example/internal/funcs"
	envtype "cryptowatcher.example/internal/types/env"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func GetEnvironmentVars(vars *envtype.EnvVars, requiredVars []string) error {

	envVars := loadEnvironmentVars()

	for _, v := range requiredVars {
		if !funcs.KeyExists(v, envVars) {
			return fmt.Errorf("Variable not found in environment file: %s\n", v)
		}

		(*vars)[v] = envVars[v]
	}

	return nil
}

func loadEnvironmentVars() map[string]string {

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

	values, err := godotenv.Read(envFile)
	if err != nil {
		fmt.Printf("Unable to open environment file. %s\n", err.Error())
		os.Exit(1)
	}

	return values
}

