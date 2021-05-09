package router

import (
	"cryptowatcher.example/internal/pkg/logga"
	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/pkg/handlers/currency"
)

func New(r *gin.Engine, l *logga.Logga) {

	c := currency.New(l)

	api := r.Group("/api")
	{
		api.GET("/currency", c.GetRecords)
	}
}