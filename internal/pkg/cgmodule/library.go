package cgmodule

import (
	"encoding/json"
	"fmt"
	"strconv"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/logga"
)

type RequestModule struct {
	ApiKey  string
	ApiHost string
	l       *logga.Logga
}

// New will return an instance of RequestModule.
func New(cfg config.CMCAggregator, logger *logga.Logga) *RequestModule {

	return &RequestModule{
		ApiHost: cfg.Host,
		ApiKey:  cfg.APIKey,
		l:       logger,
	}
}

// FetchCurrencyListing will make a request on Coin Gecko to retrieve current listings of each currency.
func (mm *RequestModule) FetchCurrencyListing(limit int) ([]CoinsID, error) {

	l := mm.l.Lg.With().Str("coingeckomodule", "FetchCurrencyListing").Logger()

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	//var dataObj CurrencyData
	var listing []CoinsID

	// https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=2&sparkline=false
	_, data, err := mm.MakeRequest("GET", "/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=2&sparkline=false", params, nil)
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
