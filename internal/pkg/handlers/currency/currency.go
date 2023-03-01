package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/models/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/handlers/pagination"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

type Handler struct {
	l  *logga.Logga
	DB dbinterface.QueryAble
}

// New - instantiate package.
func New(l *logga.Logga, db dbinterface.QueryAble) Handler {
	return Handler{
		l:  l,
		DB: db,
	}
}

// GetRecords - get record collection
func (h *Handler) GetRecords(c *gin.Context) {
	l := h.l.Lg.With().Str("currency handle", "GetRecords").Logger()
	l.Info().Msg("Request to list currency")
	cm := currency.New(h.DB, h.l)

	pg := pagination.GetPaginationVars(h.l, c)
	var count int64

	var records database.Currencies
	records, err := cm.GetRecords(pg)
	if err != nil {
		l.Error().Msgf("There was an error fetching currency records. %s", err)
		c.JSON(500, gin.H{"error": "error fetching currency records"})
	}

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": pagination.GetMetaResponse(pg, count)})
}
