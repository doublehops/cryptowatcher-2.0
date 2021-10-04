package router

import (
	"cryptowatcher.example/internal/pkg/logga"
	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/pkg/handlers/cmchistory"
	"cryptowatcher.example/internal/pkg/handlers/currency"
)

func New(r *gin.Engine, l *logga.Logga) {

	c := currency.New(l)
	ch := cmchistory.New(l)

	api := r.Group("/api")
	{
		api.GET("/currency", c.GetRecords)
		api.GET("/cmc-history/time-series/:symbol", ch.GetTimeSeriesData)
	}
}
