package aggregatorengine

import (
	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/models/currency"
	"github.com/doublehops/cryptowatcher-2.0/internal/models/history"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

type Aggregator interface {
	GetAggregatorID() uint32
	GetAggregatorName() string
	FetchLatestHistory() ([]*database.History, error)
}

type Agg struct {
	name       string
	db         dbinterface.QueryAble
	aggregator Aggregator
	logga      *logga.Logga
	currency   *currency.Model
	history    *history.Model
}

// New will instantiate a new instance of Aggregator.
func New(db dbinterface.QueryAble, agg Aggregator, logga *logga.Logga) *Agg {
	cur := currency.New(db, logga)
	hs := history.New(db, logga)

	return &Agg{
		db:         db,
		aggregator: agg,
		logga:      logga,
		currency:   cur,
		history:    hs,
	}
}

// UpdateLatestHistory will update the latest records from the aggregator.
func (a *Agg) UpdateLatestHistory() error {
	l := a.logga.Lg.With().Str("aggregatorengine", "UpdateLatestHistory").Logger()
	l.Info().Msgf("Updating history for aggregator: %s", a.name)

	curMap, err := a.currency.GetRecordsMapKeySymbol()
	if err != nil {
		return err
	}
	histories, err := a.aggregator.FetchLatestHistory()
	if err != nil {
		return err
	}

	for _, cur := range histories {
		cur.CurrencyID = a.getCurrencyID(curMap, *cur)
		_, err = a.history.CreateRecord(cur)
		if err != nil {
			l.Error().Msgf("Error adding history record: %s", cur.Symbol)
		}
	}

	return nil
}

// getCurrencyID will return the ID of the currency record in the database. It will create a new record if
// that currency doesn't already exist.
func (a *Agg) getCurrencyID(curMap map[string]uint32, h database.History) uint32 {
	l := a.logga.Lg.With().Str("aggregatorengine", "getCurrencyID").Logger()

	_, exists := curMap[h.Symbol]
	if exists {
		return curMap[h.Symbol]
	}

	// not found, create new record.
	cur := database.Currency{
		Name:   h.Name,
		Symbol: h.Symbol,
	}

	ID, err := a.currency.CreateRecord(&cur)
	if err != nil {
		l.Error().Msgf("Error adding currency: %s", h.Symbol)
	}

	return uint32(ID)
}
