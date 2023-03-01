package router

import (
	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/gin-gonic/gin"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/handlers/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/handlers/history"
)

func New(r *gin.Engine, db dbinterface.QueryAble, l *logga.Logga) {
	c := currency.New(l, db)
	ch := history.New(l, db)

	api := r.Group("/api")
	{
		api.GET("/currency", c.GetRecords)
		api.GET("/cmc-history/time-series/:symbol", ch.GetTimeSeriesData)
	}
}
