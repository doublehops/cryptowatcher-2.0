package funcs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

func GetEnvironmentVar(varName string) string {

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

	fmt.Printf("Opening environment file: %s\n", envFile)

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Printf("Unable to open environment file. %s\n", err.Error())
		os.Exit(1)
	}

	return os.Getenv(varName)
}

