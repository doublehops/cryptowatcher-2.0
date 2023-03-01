package pagination

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
)

type MetaRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Offset  int `json:"offset"`
}

type MetaResponse struct {
	Page         int   `json:"page"`
	PerPage      int   `json:"perPage"`
	TotalPages   int   `json:"totalPages"`
	TotalRecords int64 `json:"totalRecords"`
}

var (
	defaultPage    = 1
	defaultPerPage = 10
)

// GetPaginationVars - find and return pagination vars for query and meta response.
func GetPaginationVars(lg *logga.Logga, c *gin.Context) *MetaRequest {
	l := lg.Lg.With().Str("pagination", "HandlePagination").Logger()
	l.Info().Msg("Setting up pagination")

	query := c.Request.URL.Query()
	page, perPage, offset := getVars(query)

	pg := MetaRequest{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}

	return &pg
}

// getVars - Search request query for wanted var and return value, if not found, return default value.
func getVars(query map[string][]string) (int, int, int) {
	page := defaultPage
	perPage := defaultPerPage
	offset := 0

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "perPage":
			perPage, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}

	if page != 1 {
		offset = (page - 1) * perPage
	}

	return page, perPage, offset
}

// GetMetaResponse - Get pagination data to send in API response.
func GetMetaResponse(pg *MetaRequest, count int64) *MetaResponse {
	pageApprox := float64(count) / float64(pg.PerPage)
	totalPages := math.Ceil(pageApprox)
	tp := int(totalPages)

	return &MetaResponse{
		Page:         pg.Page,
		PerPage:      pg.PerPage,
		TotalPages:   tp,
		TotalRecords: count,
	}
}
