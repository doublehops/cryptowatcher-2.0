package history

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/cryptowatcher-2.0/internal/models/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/models/history"
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

// GetTimeSeriesData - get record collection
func (h *Handler) GetTimeSeriesData(c *gin.Context) {
	l := h.l.Lg.With().Str("history handle", "GetTimeSeriesData").Logger()

	symbol := c.Param("symbol")
	l.Info().Msgf("Request to retrieve time series data for symbol: %s", symbol)

	cm := currency.New(h.DB, h.l)
	chm := history.New(h.DB, h.l)

	var cur database.Currency
	err := cm.GetRecordBySymbol(&cur, symbol)
	if err != nil {
		l.Info().Msgf("Error with GetRecordBySymbol: %s", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"code": "Error with GetRecordBySymbol", "message": err.Error()})

		return
	}

	if cur.ID == 0 {
		l.Info().Msgf("symbol not found: %s", symbol)
		c.JSON(http.StatusNotFound, gin.H{"code": "symbol not found", "message": "Symbol not found"})

		return
	}

	searchParams, err := h.getSearchParams(c)
	if err != nil {
		l.Info().Msgf("Error processing request for symbol: %s; error: %s", symbol, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": "Error processing request", "message": err.Error()})

		return
	}

	records, err := chm.GetPriceTimeSeriesData(symbol, searchParams)
	if err != nil {
		l.Info().Msgf("Error with GetPriceTimeSeriesData: %s", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"code": "Error with GetPriceTimeSeriesData", "message": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": records})
}

// getSearchParams - get search parameters to fetch records by.
func (h *Handler) getSearchParams(c *gin.Context) (*history.SearchParams, error) {
	l := h.l.Lg.With().Str("history handle", "getSearchParams").Logger()

	var t string
	var params history.SearchParams

	now := time.Now()
	secs := now.Unix()

	// Get timeFrom
	t, _ = c.GetQuery("timeFrom")
	if t != "" {
		ct, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return &params, err
		}
		params.TimeFromUnix = ct
	} else {
		params.TimeFromUnix = secs - 60*60*24*7 // 7 days ago
	}

	// Get timeTo
	t, _ = c.GetQuery("timeTo")
	if t != "" {
		ct, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			return &params, err
		}
		params.TimeToUnix = ct
	} else {
		params.TimeToUnix = secs
	}

	if params.TimeFrom > params.TimeTo {
		return &params, fmt.Errorf("time-to cannot be earlier than time-fome")
	}

	// Convert times to strings - 2006-01-02 15:04:05
	tf := time.Unix(params.TimeFromUnix, 0)
	params.TimeFrom = tf.Format("2006-01-02 15:04:05")
	tt := time.Unix(params.TimeToUnix, 0)
	params.TimeTo = tt.Format("2006-01-02 15:04:05")

	l.Debug().Msg("Search params calculated")
	l.Debug().Msgf("%v", params)

	return &params, nil
}
