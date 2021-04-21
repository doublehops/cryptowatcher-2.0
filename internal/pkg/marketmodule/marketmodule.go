package marketmodule

import (
	"encoding/json"
	"strconv"

	"gorm.io/gorm"

	"cryptowatcher.example/internal/models/cmchistory"
	"cryptowatcher.example/internal/models/coin"
	"cryptowatcher.example/internal/pkg/cmchttp"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/api"
	"cryptowatcher.example/internal/types/database"
)

type marketmodule struct {
	db     *gorm.DB
	ApiKey string
	cmcr   cmchttp.Requester
	l      *logga.Logga
}

func New(db *gorm.DB, ApiKey string, logger *logga.Logga) *marketmodule {

	return &marketmodule{
		db:     db,
		ApiKey: ApiKey,
		l:      logger,
	}
}

func (mm *marketmodule) SaveCurrencyListing(numberToRetrieve int) (string, error) {

	l := mm.l.Lg.With().Str("marketmodule", "SaveCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")

	tx := mm.db.Begin()

	cm := coin.New(tx, mm.l)

	currencies, _ := mm.GetCurrencyListing(numberToRetrieve)

	for _, c := range currencies.Currencies {

		cr := cm.GetCoinBySymbol(c.Symbol)

		if cr.ID == 0 {
			crNew := database.Coin{
				Name:   c.Name,
				Symbol: c.Symbol,
			}

			result := cm.CreateCoin(&crNew)
			if result.Error != nil {
				l.Error().Msg("Error saving coin record to database")
				l.Error().Msgf("%v", result.Error)
				return "", result.Error
			}
		}

		// @TODO: Add record to CmcHistory

		cmcm := cmchistory.New(mm.db, mm.l)

		cmcr := &database.CmcHistory{
			Name:              c.Name,
			Symbol:            c.Symbol,
			Slug:              c.Slug,
			NumMarketPairs:    c.NumMarketPairs,
			DateAdded:         c.DateAdded,
			MaxSupply:         c.MaxSupply,
			CirculatingSupply: c.CirculatingSupply,
			TotalSupply:       c.TotalSupply,
			CmcRank:           c.CmcRank,
			QuotePrice:        c.Quote.USDObj.Price,
			Volume24h:         c.Quote.USDObj.Volume24Hours,
			PercentChange1h:   c.Quote.USDObj.PercentChange1Hour,
			PercentChange24h:  c.Quote.USDObj.PercentChange24Hours,
			PercentChange7D:   c.Quote.USDObj.PercentChange7Days,
			PercentChange30D:  c.Quote.USDObj.PercentChange30Days,
			PercentChange60D:  c.Quote.USDObj.PercentChange60Days,
			PercentChange90D:  c.Quote.USDObj.PercentChange90Days,
			MarketCap:         c.Quote.USDObj.MarketCap,
		}

		cmcm.CreateRecord(cmcr)
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

	cmcr := cmchttp.New(mm.ApiKey, mm.l)

	_, data, err := cmcr.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating marketmodule request client")
		l.Error().Msg(err.Error())
		return nil, err
	}

	var listing types.CurrencyListing
	err = json.Unmarshal(data, &listing)
	if err != nil {
		l.Error().Msg("There was an error unmarshalling json marketmodule response")
		l.Error().Msg(err.Error())
		return nil, err
	}

	l.Info().Msgf("%d currencies returned\n\n", len(listing.Currencies))

	return &listing, nil
}
