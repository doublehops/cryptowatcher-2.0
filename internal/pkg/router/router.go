package router

import (
	"cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/pkg/logga"
	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/pkg/handlers/cmchistory"
	"cryptowatcher.example/internal/pkg/handlers/currency"
)

func New(r *gin.Engine, db dbinterface.QueryAble, l *logga.Logga) {

	c := currency.New(l, db)
	ch := cmchistory.New(l, db)

	api := r.Group("/api")
	{
		api.GET("/currency", c.GetRecords)
		api.GET("/cmc-history/time-series/:symbol", ch.GetTimeSeriesData)
	}
}
