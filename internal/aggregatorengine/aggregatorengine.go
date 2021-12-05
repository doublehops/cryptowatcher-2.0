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
	FetchLatestHistory() (database.Histories, error)
}

type Agg struct {
	db       dbinterface.QueryAble
	logga    *logga.Logga
	currency *currency.Model
	history *history.Model
}

func New(db dbinterface.QueryAble, logga *logga.Logga) *Agg {

	cur := currency.New(db, logga)
	history := history.New(db, logga)
	return &Agg{
		db:       db,
		logga:    logga,
		currency: cur,
		history: history,
	}
}

func (a *Agg) UpdateLatestHistory(agg Aggregator) error {

	l := a.logga.Lg.With().Str("aggregator", "UpdateLatestHistory").Logger()

	curMap, err := a.getCurrencyMapping()
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

func (a *Agg) getCurrencyID(curMap map[string]uint32, h database.History) uint32 {

	l := a.logga.Lg.With().Str("aggregator", "getCurrencyID").Logger()

	var curID uint32

	_, exists := curMap[h.Symbol]
	if exists {
		return curID
	}

	// not found, create new record.
	var cur database.Currency
	cur.Name = h.Name
	cur.Symbol = h.Symbol

	ID, err := a.currency.CreateRecord(&cur)
	if err != nil {
		l.Error().Msgf("Error adding currency: %s", cur.Symbol)
	}

	curID = uint32(ID)

	return curID
}

func (a *Agg) getCurrencyMapping() (map[string]uint32, error) {

	return a.currency.GetRecordsMapKeySymbol()
}
