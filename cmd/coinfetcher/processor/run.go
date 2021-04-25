package processor

import (
	"cryptowatcher.example/internal/env"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
	"github.com/davecgh/go-spew/spew"
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
		ec := cm.GetCoinBySymbol(c.Symbol)
		spew.Dump(ec)
		if ec.ID > 0 { // Coin already in database.
			continue
		}

		crNew := database.Currency{
			Name:   c.Name,
			Symbol: c.Symbol,
		}

		cm.CreateCurrency(&crNew)
	}

	return nil
}