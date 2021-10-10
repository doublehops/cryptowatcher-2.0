package currency

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/db"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"cryptowatcher.example/internal/pkg/logga"
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

// GetRecords - get record collection
func (h *Handler) GetRecords(c *gin.Context) {

	l := h.l.Lg.With().Str("currency handle", "GetRecords").Logger()
	l.Info().Msg("Request to list currency")

	// Setup db connection.
	//db := orm.Connect(h.l, h.e)
	DB, err := db.New(h.l, h.e)
	if err!= nil {
		// @todo handle error
	}
	cm := currency.New(DB, h.l)

	pg := pagination.GetPaginationVars(h.l, c)
	var count int64

	var records database.Currencies
	cm.GetRecords(&records, pg, &count)

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": pagination.GetMetaResponse(pg, count)})
}
