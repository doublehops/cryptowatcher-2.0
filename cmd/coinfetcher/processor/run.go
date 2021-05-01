package processor

import (
	"gorm.io/gorm"

	"cryptowatcher.example/internal/models/cmchistory"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/env"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Runner struct {
	e    *env.Env
	l    *logga.Logga
	db   *gorm.DB
	cmcm *cmcmodule.CmcModule
}

func New(e *env.Env, l *logga.Logga, db *gorm.DB, cmcm *cmcmodule.CmcModule) *Runner {

	return &Runner{
		e:    e,
		l:    l,
		db:   db,
		cmcm: cmcm,
	}
}

func (r *Runner) Run() error {

	l := r.l.Lg.With().Str("main", "Run").Logger()
	l.Info().Msg("Running currency fetcher")

	currencies, err := r.cmcm.FetchCurrencyListing(20)
	if err != nil {
		r.l.Error("Unable to get currency listing from CMC module")
		return err
	}

	cm := currency.New(r.db, r.l)
	cmch := cmchistory.New(r.db, r.l)

	// Add coins if not already in the database.
	for _, c := range currencies {

		var cur database.Currency

		cm.GetRecordBySymbol(&cur, c.Symbol)
		if cur.ID == 0 { // Currency not yet in database.

			cur.Name = c.Name
			cur.Symbol = c.Symbol

			err := cm.CreateRecord(&cur)
			if err != nil {
				l.Error().Msgf("Error adding currency: %s", cur.Symbol)
			}
		}

		cmcr := &database.CmcHistory{
			Name:              c.Name,
			Currency:          cur,
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

		err = cmch.CreateRecord(cmcr)
		if err != nil {
			l.Error().Msgf("Error adding currency: %s", cur.Symbol)
		}
	}

	return nil
}
