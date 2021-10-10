package processor

import (
	"cryptowatcher.example/internal/dbinterface"
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
	db   dbinterface.QueryAble
	cmcm *cmcmodule.CmcModule
}

func New(e *env.Env, l *logga.Logga, db dbinterface.QueryAble, cmcm *cmcmodule.CmcModule) *Runner {

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

	for _, c := range currencies {

		var cur database.Currency

		var curID uint32
		curMap := make(map[string]uint32)
		cm.GetRecordsMapKeySymbol(&curMap)

		// Check if currency already exists in db.
		_, exists := curMap[c.Symbol]

		if !exists {

			cur.Name = c.Name
			cur.Symbol = c.Symbol

			_, err = cm.CreateRecord(&cur)
			if err != nil {
				l.Error().Msgf("Error adding currency: %s", cur.Symbol)
			}
			curID = cur.ID
		} else {
			curID = curMap[c.Symbol]
		}

		cmcr := &database.CmcHistory{
			Name:              c.Name,
			CurrencyID:        curID,
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

		_, err = cmch.CreateRecord(cmcr)
		if err != nil {
			l.Error().Msgf("Error adding currency: %s", cur.Symbol)
		}
	}

	return nil
}
