package cmchistory

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"os"
	"strconv"
	"time"

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

	searchParams, err := h.getSearchParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "Error processing request", "message": err.Error()})
		return
	}

	pg := pagination.GetPaginationVars(h.l, c)
	var count int64

	var records database.CmcHistories
	chm.GetTimeSeriesData(symbol, searchParams, &records, pg, &count)

	c.JSON(http.StatusOK, gin.H{"data": records, "meta": pagination.GetMetaResponse(pg, count)})
}

// getSearchParams - get search parameters to fetch records by.
func (h *Handler) getSearchParams(c *gin.Context) (*cmchistory.SearchParams, error) {

	l := h.l.Lg.With().Str("cmchistory handle", "getSearchParams").Logger()

	var t string
	var params cmchistory.SearchParams

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

	// Get interval - try to keep the number of intervals to around 150 for the given time period,
	// to not overwork frontend graph library and browser.
	var day int64 = 60 * 60 * 24
	var week = day * 7
	var month = day * 30
	var year = day * 365

	i := c.Param("interval")
	if i != "" {
		it, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return &params, err
		}
		params.Interval = it
	} else {
		duration := params.TimeToUnix - params.TimeFromUnix

		if duration >= year {
			params.Interval = day * 2
		} else if duration >= month {
			params.Interval = 60 * 60 * 8 // every 8 hours.
		} else if duration >= week {
			params.Interval = 60 * 60 // every hour.
		} else {
			params.Interval = 60 * 10
		}
	}

	// Convert times to strings - 2006-01-02 15:04:05
	tf := time.Unix(params.TimeFromUnix, 0)
	params.TimeFrom = tf.Format("2006-01-02 15:04:05")
	tt := time.Unix(params.TimeToUnix, 0)
	params.TimeTo = tt.Format("2006-01-02 15:04:05")

	spew.Dump(params)

	return &params, nil
}
