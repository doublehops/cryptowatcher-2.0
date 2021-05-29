package cmchistory

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/models/cmchistory"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
)

type Handler struct {
	l *logga.Logga
	e *env.Env
}

// New - instantiate package.
func New(l *logga.Logga) Handler {

	// Setup environment.
	e, err := env.New(l)
	if err != nil {
		l.Lg.Error().Msg(err.Error())
		os.Exit(1)
	}
	return Handler{
		l: l,
		e: e,
	}
}

// GetTimeSeriesData - get record collection
func (h *Handler) GetTimeSeriesData(c *gin.Context) {

	l := h.l.Lg.With().Str("cmchistory handle", "GetTimeSeriesData").Logger()

	symbol := c.Param("symbol")
	l.Info().Msgf("Request to retrieve time series data for symbol: %s", symbol)

	// Setup database connection.
	db := orm.Connect(h.l, h.e)
	cm := currency.New(db, h.l)
	chm := cmchistory.New(db, h.l)

	var cur database.Currency
	cm.GetRecordBySymbol(&cur, symbol)
	
	if cur.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": "symbol not found", "message": "Symbol not found"})
		return
	}

	pg := pagination.GetPaginationVars(h.l, c)
	var count int64

	var records database.CmcHistories
	chm.GetTimeSeriesData(symbol, &records, pg, &count)

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": pagination.GetMetaResponse(pg, count)})
}
