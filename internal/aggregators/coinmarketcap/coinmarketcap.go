package coinmarketcap

import (
	"cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/pkg/cmcmodule"
	"cryptowatcher.example/internal/pkg/config"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

const (
	aggregatorID uint32 = 1
	aggregatorName string = "cmc"
)


type Runner struct {
	cfg  config.CMCAggregator
	l    *logga.Logga
	db   dbinterface.QueryAble
	cmcm *cmcmodule.CmcModule
}

func New(cfg config.CMCAggregator, l *logga.Logga, db dbinterface.QueryAble, cmcm *cmcmodule.CmcModule) *Runner {

	return &Runner{
		cfg:  cfg,
		l:    l,
		db:   db,
		cmcm: cmcm,
	}
}

func (r *Runner) FetchLatestHistory() (database.Histories, error) {
	var histories database.Histories

	l := r.l.Lg.With().Str("main", "Run").Logger()
	l.Info().Msg("Running currency fetcher")

	currencies, err := r.cmcm.FetchCurrencyListing(20)
	if err != nil {
		r.l.Error("Unable to get currency listing from CMC module")
		return histories, err
	}

	//cm := currency.New(r.db, r.l)

	for _, c := range currencies {

		//var cur database.Currency
		//
		//var curID uint32
		//curMap := cm.GetRecordsMapKeySymbol()
		//
		//// Check if currency already exists in the database.
		//_, exists := curMap[c.Symbol]
		////
		//if !exists {
		//
		//	cur.Name = c.Name
		//	cur.Symbol = c.Symbol
		//
		//	_, err = cm.CreateRecord(&cur)
		//	if err != nil {
		//		l.Error().Msgf("Error adding currency: %s", cur.Symbol)
		//	}
		//	curID = cur.ID
		//} else {
		//	curID = curMap[c.Symbol]
		//}

		history := &database.History{
			AggregatorID:      r.GetAggregatorID(),
			//CurrencyID:        curID,
			Name:              c.Name,
			Symbol:            c.Symbol,
			Slug:              c.Slug,
			NumMarketPairs:    c.NumMarketPairs,
			DateAdded:         c.DateAdded,
			MaxSupply:         c.MaxSupply,
			CirculatingSupply: c.CirculatingSupply,
			TotalSupply:       c.TotalSupply,
			Rank:           c.CmcRank,
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

		histories = append(histories, history)
		//_, err = cmch.CreateRecord(cmcr)
		//if err != nil {
		//	l.Error().Msgf("Error adding currency: %s", cur.Symbol)
		//}
	}

	return histories, nil
}

// GetAggregatorName will return the aggregator name for log messages
func (r *Runner) GetAggregatorName() string {
	return aggregatorName
}

// GetAggregatorID will return the aggregator ID to distinguish which aggregator each currency record came from.
func (r *Runner) GetAggregatorID() uint32 {
	return aggregatorID
}
