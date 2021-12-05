package aggregatorengine

import (
	"cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/models/currency"
	"cryptowatcher.example/internal/models/history"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Aggregator interface {
	GetAggregatorID() uint32
	GetAggregatorName() string
	FetchLatestHistory() (database.Histories, error)
}

type Agg struct {
	name 	 string
	db       dbinterface.QueryAble
	logga    *logga.Logga
	currency *currency.Model
	history  *history.Model
}

// New will instantiate a new instance of Aggregator.
func New(db dbinterface.QueryAble, logga *logga.Logga) *Agg {

	cur := currency.New(db, logga)
	hs := history.New(db, logga)

	return &Agg{
		db:       db,
		logga:    logga,
		currency: cur,
		history:  hs,
	}
}

// UpdateLatestHistory will update latest records from the aggregator.
func (a *Agg) UpdateLatestHistory(agg Aggregator) error {

	l := a.logga.Lg.With().Str("aggregatorengine", "UpdateLatestHistory").Logger()
	l.Info().Msgf("Updating history for aggregator: %s", a.name)

	curMap, err := a.currency.GetRecordsMapKeySymbol()
	if err != nil {
		return err
	}
	histories, err := agg.FetchLatestHistory()
	if err != nil {
		return err
	}

	for _, cur := range histories {
		cur.ID = a.getCurrencyID(curMap, *cur)
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
		Name: h.Name,
		Symbol: h.Symbol,
	}

	ID, err := a.currency.CreateRecord(&cur)
	if err != nil {
		l.Error().Msgf("Error adding currency: %s", h.Symbol)
	}

	return uint32(ID)
}