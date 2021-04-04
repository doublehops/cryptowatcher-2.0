package marketmodule

import (
	"cryptowatcher.example/internal/models/coin"
	"cryptowatcher.example/internal/types/database"
	"encoding/json"
	"gorm.io/gorm"
	"strconv"

	"cryptowatcher.example/internal/pkg/cmchttp"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/api"
)

type marketmodule struct {
	db *gorm.DB
	ApiKey string
	cmcr   cmchttp.Requester
	l *logga.Logga
}

func New(db *gorm.DB, ApiKey string, logger *logga.Logga) *marketmodule {

	return &marketmodule{
		db: db,
		ApiKey: ApiKey,
		l: logger,
	}
}

func (mm *marketmodule) SaveCurrencyListing(numberToRetrieve int) (string, error) {

	l := mm.l.Lg.With().Str("marketmodule", "SaveCurrencyListing").Logger()

	l.Info().Msg("---  Fetching currencies  ---")

	cm := coin.New(mm.db, mm.l)

	currencies, _ := mm.GetCurrencyListing(numberToRetrieve)

	for _, c := range currencies.Currencies {

		cr, err := cm.GetCoinBySymbol(c.Symbol)
		if err != nil {
			return "", err
		}

		l.Info().Msgf(">>>>  ID found for `%s` is: %v", c.Symbol, cr.ID)

		if cr.ID == 0 {
			crNew := database.Coin{
				Name: c.Name,
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
