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

func (m *Model) GetCoinBySymbol(s string) *database.Currency {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	r := database.Currency{}
	m.db.Find(&r, "symbol = ?", s)

	return &r
}

func (m *Model) CreateCurrency(r *database.Currency) *gorm.DB {

	l := m.l.Lg.With().Str("currency", "CreateCurrency").Logger()
	l.Info().Msgf("Adding currency: %s", r.Symbol)

	result := m.db.Create(&r)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
		return result
	}

	return result
}
