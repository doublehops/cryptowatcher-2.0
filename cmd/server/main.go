package main

import (
	"cryptowatcher.example/internal/pkg/logga"
	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/pkg/router"
)

func main() {

	// Setup logger.
	l := logga.New()

	r := gin.Default()
	router.New(r, l)

	r.Run(":8080")
}
