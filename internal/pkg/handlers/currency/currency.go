package currency

import (
	"cryptowatcher.example/internal/dbinterface"
	"github.com/gin-gonic/gin"
	"net/http"

	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Handler struct {
	l   *logga.Logga
	DB dbinterface.QueryAble
}

// New - instantiate package.
func New(l *logga.Logga, db dbinterface.QueryAble) Handler {

	return Handler{
		l: l,
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
	cm.GetRecords(&records, pg, &count)

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": pagination.GetMetaResponse(pg, count)})
}
