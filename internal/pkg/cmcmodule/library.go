package cmcmodule

import (
	"encoding/json"
	"strconv"

	"cryptowatcher.example/internal/pkg/logga"
)

type CmcModule struct {
	ApiKey  string
	ApiHost string
	l       *logga.Logga
}

func New(ApiKey, ApiHost string, logger *logga.Logga) *CmcModule {

	return &CmcModule{
		ApiKey:  ApiKey,
		ApiHost: ApiHost,
		l:       logger,
	}
}

func (mm *CmcModule) FetchCurrencyListing(limit int) ([]*Currency, error) {

	l := mm.l.Lg.With().Str("marketmodule", "GetCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	var dataObj Data
	var listing []*Currency

	_, data, err := mm.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating marketmodule request client")
		l.Error().Msg(err.Error())
		return listing, err
	}

	err = json.Unmarshal(data, &dataObj)
	if err != nil {
		l.Error().Msg("There was an error unmarshalling json marketmodule response")
		l.Error().Msg(err.Error())
		l.Error().Msgf("%#v", err)
		return listing, err
	}

	listing = dataObj.Currencies

	l.Info().Msgf("%d currencies returned", len(listing))

	return listing, nil
}
