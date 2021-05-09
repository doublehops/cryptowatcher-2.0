package currency

import (
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"cryptowatcher.example/internal/pkg/orm"
	"cryptowatcher.example/internal/types/database"
	"net/http"
	"os"

	"cryptowatcher.example/internal/pkg/logga"
	"github.com/gin-gonic/gin"
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

	l := h.l.Lg.With().Str("currency", "ListRecords").Logger()
	l.Info().Msg("Request to list currency")

	// Setup database connection.
	db := orm.Connect(h.l, h.e)
	cm := currency.New(db, h.l)

	pg := pagination.GetPaginationVars(h.l, c)

	var records database.Currencies
	cm.GetRecords(&records, pg)

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": "meta-here"})
}
