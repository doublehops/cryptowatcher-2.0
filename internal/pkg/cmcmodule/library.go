package cmcmodule

import (
	"encoding/json"
	"fmt"
	"strconv"

	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/logga"
)

type CmcModule struct {
	ApiKey  string
	ApiHost string
	l       *logga.Logga
}

func New(cfg config.CMCAggregator, logger *logga.Logga) *CmcModule {

	return &CmcModule{
		ApiHost: cfg.Host,
		ApiKey:  cfg.APIKey,
		l:       logger,
	}
}

func (mm *CmcModule) FetchCurrencyListing(limit int) ([]*Currency, error) {

	l := mm.l.Lg.With().Str("marketmodule", "FetchCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	var dataObj CurrencyData
	var listing []*Currency

	_, data, err := mm.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		return listing, fmt.Errorf("\"There was an error instantiating marketmodule request client. %w", err)
	}

	err = json.Unmarshal(data, &dataObj)
	if err != nil {
		return listing, fmt.Errorf("there was an error unmarshalling json marketmodule response. %w", err)
	}

	listing = dataObj.Currencies

	l.Info().Msgf("%d currencies returned", len(listing))

	return listing, nil
}
