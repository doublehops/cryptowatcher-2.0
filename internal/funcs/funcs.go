package funcs

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetEnvironmentVar(varName string) string {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to open environment file")
		os.Exit(1)
	}

	return os.Getenv(varName)
}

