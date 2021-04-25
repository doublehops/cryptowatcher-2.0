package currency

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

func (m *Model) GetCoinBySymbol(record *database.Currency, s string) {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	m.db.Find(&record, "symbol = ?", s)
}

func (m *Model) CreateCurrency(record *database.Currency) (error) {

	l := m.l.Lg.With().Str("currency", "CreateCurrency").Logger()
	l.Info().Msgf("Adding currency: %s", record.Symbol)

	result := m.db.Create(&record)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
		return result.Error
	}

	return nil
}
