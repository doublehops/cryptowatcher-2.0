package coin

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

func (m *Model) GetCoinBySymbol(s string) *database.Coin {

	l := m.l.Lg.With().Str("coin", "GetCoinBySymbol").Logger()

	l.Info().Msgf("Fetching coin by symbol: %s", s)

	r := database.Coin{}
	m.db.Find(&r, "symbol = ?", s)

	return &r
}

func (m *Model) CreateCoin(r *database.Coin) *gorm.DB {

	l := m.l.Lg.With().Str("coin", "CreateCoin").Logger()

	l.Info().Msgf("Adding coin: %s", r.Symbol)

	result := m.db.Create(&r)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
		return result
	}

	return result
}
