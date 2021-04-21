package processor

import (
	"cryptowatcher.example/internal/env"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/logga"
	"github.com/davecgh/go-spew/spew"
)

type Runner struct {
	e    *env.Env
	l    *logga.Logga
	cmcm *cmcmodule.CmcModule
}

func New(e *env.Env, l *logga.Logga, cmcm *cmcmodule.CmcModule) *Runner {

	return &Runner{
		e:    e,
		l:    l,
		cmcm: cmcm,
	}
}

func (r *Runner) Run() error {

	l := r.l.Lg.With().Str("main", "Run").Logger()
	l.Info().Msg("Running coin fetcher")

	coins, err := r.cmcm.GetCurrencyListing(20)
	if err != nil {
		r.l.Error("Unable to get currency listing from CMC module")
		return err
	}

	spew.Dump(coins)

	for _, coin := range coins {

	}

	return nil
}


//func Run(e *env.Env, l *logga.Logga, cmcm *cmcmodule.CmcModule) error {
//
//	//l := l.Lg.With().Str("main", "Run").Logger()
//	//l.Info().Msg("Running coin fetcher")
//
//	coins, err := cmcm.GetCurrencyListing(20)
//	if err != nil {
//		l.Error("Unable to get currency listing from CMC module")
//		return err
//	}
//
//	spew.Dump(coins)
//
//	//for _, coin := range coins {
//	//
//	//}
//	return nil
//}
