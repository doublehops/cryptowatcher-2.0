package cmchistory

import (
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db *gorm.DB
	l  *logga.Logga
}

type SearchParams struct {
	TimeFrom     string
	TimeTo       string
	TimeFromUnix int64
	TimeToUnix   int64
}

// New - creates new instance of cmchistory.
func New(db *gorm.DB, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l:  logger,
	}
}

// CreateRecord inserts a new record into the cmc_history table.
func (m *Model) CreateRecord(record *database.CmcHistory) error {

	l := m.l.Lg.With().Str("cmchistory", "CreateRecord").Logger()
	l.Info().Msgf("Adding cmc record: %s", record.Symbol)

	result := m.db.Create(&record)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to db. %v", result.Error)
		return result.Error
	}

	return nil
}

// GetRecordByID will return the requested record from the db by its ID.
func (m *Model) GetRecordByID(record *database.CmcHistory, ID uint32) error {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordByID").Logger()
	l.Info().Msgf("Retrieving cmchistory record by ID: %d", ID)

	m.db.First(&record, "id = ?", ID)

	return nil
}

// GetPriceTimeSeriesData will return records grouped together in X number of groups with `quote_price` averaged out per group.
func (m *Model) GetPriceTimeSeriesData(symbol string, searchParams *SearchParams, records *database.CmcHistoriesPriceTimeSeriesData) {

	l := m.l.Lg.With().Str("cmchistory", "GetTimeSeriesData").Logger()
	l.Info().Msgf("Fetching cmchistory records for symbol: %s", symbol)

	buckets := 5 // number of buckets to group the records in and determine average for.

	m.db.Raw(TimeSeriesSlicedPeriodQuery, buckets, symbol, searchParams.TimeFrom, searchParams.TimeTo).Scan(records)
}

// GetRecordsBySymbol will return a collection of CmcHistory records from the db.
func (m *Model) GetRecordsBySymbol(symbol string) ([]database.CmcHistory, error) {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordsBySymbol").Logger()
	l.Info().Msgf("Retrieving cmchistory record by symbol: %s", symbol)

	var records []database.CmcHistory

	m.db.Find(&records, "symbol = ?", symbol)

	return records, nil
}
