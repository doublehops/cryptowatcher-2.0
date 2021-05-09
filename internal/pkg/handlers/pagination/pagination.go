package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"cryptowatcher.example/internal/pkg/logga"
)

type Meta struct {
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
	Offset  int    `json:"offset"`
}

var defaultPage = 1
var defaultPerPage = 10


// GetPaginationVars - find and return pagination vars for query and meta response.
func GetPaginationVars(lg *logga.Logga, c *gin.Context) *Meta {

	l := lg.Lg.With().Str("pagination", "HandlePagination").Logger()
	l.Info().Msg("Setting up pagination")

	query := c.Request.URL.Query()
	page, perPage, offset := getPaginationVars(query)

	pg := Meta{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
	}

	return &pg
}

// getVar - Search request query for wanted var and return value, if not found, return default value.
func getPaginationVars(query map[string][]string) (int, int, int) {

	page := defaultPage
	perPage := defaultPerPage
	var offset = 0

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "perPage":
			perPage, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}

	if page != 1 {
		offset = (page - 1) * perPage
	}

	return page, perPage, offset
}
