package cmchistory

import (
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"gorm.io/gorm"

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
	Interval     int64
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
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
		return result.Error
	}

	return nil
}

// GetRecordByID will return the requested record from the database by its ID.
func (m *Model) GetRecordByID(record *database.CmcHistory, ID uint32) error {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordByID").Logger()
	l.Info().Msgf("Retrieving cmchistory record by ID: %d", ID)

	m.db.First(&record, "id = ?", ID)

	return nil
}

// GetTimeSeriesData will return model records.
func (m *Model) GetTimeSeriesData(symbol string, searchParams *SearchParams, records *database.CmcHistories, pg *pagination.MetaRequest, count *int64) {

	l := m.l.Lg.With().Str("cmchistory", "GetTimeSeriesData").Logger()
	l.Info().Msgf("Fetching cmchistory records")

	m.db.Find(records).
		Where("symbol", symbol).
		Where("created_at > ?", searchParams.TimeFrom).
		Where("created_at < ?", searchParams.TimeTo).
		Count(count)

	m.db.Where("symbol", symbol).Debug().
		Where("created_at > ?", searchParams.TimeFrom).
		Where("created_at < ?", searchParams.TimeTo).
		Find(records)
}

// GetRecordsBySymbol will return a collection of CmcHistory records from the database.
func (m *Model) GetRecordsBySymbol(symbol string) ([]database.CmcHistory, error) {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordsBySymbol").Logger()
	l.Info().Msgf("Retrieving cmchistory record by symbol: %s", symbol)

	var records []database.CmcHistory

	m.db.Find(&records, "symbol = ?", symbol)

	return records, nil
}
