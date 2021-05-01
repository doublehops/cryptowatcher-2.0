package cmchistory

import (
	"gorm.io/gorm"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db *gorm.DB
	l *logga.Logga
}

func New(db *gorm.DB, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l: logger,
	}
}

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

func (m *Model) GetRecordByID(record *database.CmcHistory, ID uint32) error {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordByID").Logger()
	l.Info().Msgf("Retrieving cmchistory record by ID: %d", ID)

	m.db.First(&record, "id = ?", ID)

	return nil
}

func (m *Model) GetRecordsBySymbol(symbol string) ([]database.CmcHistory, error) {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordsBySymbol").Logger()
	l.Info().Msgf("Retrieving cmchistory record by symbol: %s", symbol)

	var records []database.CmcHistory

	m.db.Find(&records, "symbol = ?", symbol)

	return records, nil
}
