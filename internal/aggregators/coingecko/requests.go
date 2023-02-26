package coingecko

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FetchCurrencyListing will make a request on Coingecko to retrieve current listings of each currency.
func (r *Runner) FetchCurrencyListing(limit int) ([]*Currency, error) {

	l := r.l.Lg.With().Str(packageName, "FetchCurrencyListing").Logger()

	params := map[string]string{
		"per_page":    strconv.Itoa(limit),
		"sparklines":  "false",
		"order":       "market_cap_desc",
		"vs_currency": "usd",
	}

	var listing []*Currency

	_, data, err := r.MakeRequest("GET", "/api/v3/coins/markets", params, nil)
	if err != nil {
		errMsg := fmt.Errorf("there was an error instantiating marketmodule request client. %w", err)
		l.Error().Msg(errMsg.Error())
		return listing, errMsg
	}

	err = json.Unmarshal(data, &listing)
	if err != nil {
		errMsg := fmt.Errorf("there was an error unmarshalling json marketmodule response. %w", err)
		l.Error().Msg(errMsg.Error())

		return listing, errMsg
	}

	l.Info().Msgf("%d currencies returned", len(listing))

	return listing, nil
}
