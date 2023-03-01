package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FetchCurrencyListing will make a request on coinmarketcap to retrieve current listings of each currency.
func (r *Runner) FetchCurrencyListing(limit int) ([]*Currency, error) {
	l := r.l.Lg.With().Str(packageName, "FetchCurrencyListing").Logger()

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	var dataObj CurrencyData
	var listing []*Currency

	_, data, err := r.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		errMsg := fmt.Errorf("there was an error instantiating marketmodule request client. %w", err)
		l.Error().Msg(errMsg.Error())

		return listing, errMsg
	}

	err = json.Unmarshal(data, &dataObj)
	if err != nil {
		errMsg := fmt.Errorf("there was an error unmarshalling json marketmodule response. %w", err)
		l.Error().Msg(errMsg.Error())

		return listing, errMsg
	}

	listing = dataObj.Currencies

	l.Info().Msgf("%d currencies returned", len(listing))

	return listing, nil
}
