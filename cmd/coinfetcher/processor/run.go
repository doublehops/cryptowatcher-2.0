package processor

import (
	"cryptowatcher.example/internal/env"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
	"gorm.io/gorm"
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

	// Add coins if not already in the database.
	for _, c := range currencies {

		var cur database.Currency

		cm.GetCoinBySymbol(&cur, c.Symbol)
		if cur.ID == 0 { // Currency not yet in database.

			cur.Name = c.Name
			cur.Symbol = c.Symbol

			err := cm.CreateCurrency(&cur)
			if err != nil {
				l.Error().Msgf("Error adding currency: %s", cur.Symbol)
			}
		}
	}

	return nil
}