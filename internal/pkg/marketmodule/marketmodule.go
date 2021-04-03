package marketmodule

import (
	"encoding/json"
	"strconv"

	"cryptowatcher.example/internal/pkg/cmchttp"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types"
)

type marketmodule struct {
	ApiKey string
	cmcr   cmchttp.Requester
	l *logga.Logga
}

func New(ApiKey string, logger *logga.Logga) *marketmodule {

	return &marketmodule{
		ApiKey: ApiKey,
		l: logger,
	}
}

func (mm *marketmodule) SaveCurrencyListing(numberToRetrieve int) (string, error) {

	l := mm.l.Lg.With().Str("marketmodule", "SaveCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")
	l.Info().Msgf("Fetching top %d but dispalying top %d\n\n", numberToRetrieve)

	currencies, _ := mm.GetCurrencyListing(numberToRetrieve)

	for _, c := range currencies.Currencies {

		ratio := c.Quote.USDObj.MarketCap / c.Quote.USDObj.Volume24Hours

		_ = types.CurrencySortBase{
			Name:           c.Name,
			Rank:           c.CmcRank,
			Symbol:         c.Symbol,
			MarketCap:      c.Quote.USDObj.MarketCap,
			Volume24h:      c.Quote.USDObj.Volume24Hours,
			CapVolumeRatio: ratio,
		}

		// @TODO: Save to database
	}

	return "", nil
}

func (mm *marketmodule) GetCurrencyListing(limit int) (*types.CurrencyListing, error) {

	l := mm.l.Lg.With().Str("marketmodule", "GetCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	cmcr := cmchttp.New(mm.ApiKey)

	_, data, err := cmcr.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating mm request client")
		l.Error().Msg(err.Error())
		return nil, err
	}

	var listing types.CurrencyListing
	err = json.Unmarshal(data, &listing)
	if err != nil {
		l.Error().Msg("There was an error unmarshalling json mm response")
		l.Error().Msg(err.Error())
		return nil, err
	}

	l.Info().Msgf("%d currencies returned\n\n", len(listing.Currencies))

	return &listing, nil
}
