package coin

import (
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
	"gorm.io/gorm"
)

type Coin struct {
	db *gorm.DB
	l *logga.Logga
}

func New(db *gorm.DB, logger *logga.Logga) *Coin {

	return &Coin{
		db: db,
		l: logger,
	}
}

func (m *Coin) GetCoinBySymbol(s string) (*database.Coin, error) {

	l := m.l.Lg.With().Str("coin", "GetCoinBySymbol").Logger()

	l.Info().Msgf("Fetching coin by symbol: %s", s)

	r := database.Coin{}
	m.db.Find(&r, "symbol = ?", s)

	return &r, nil
}

func (m *Coin) CreateCoin(r *database.Coin) *gorm.DB {

	l := m.l.Lg.With().Str("coin", "CreateCoin").Logger()

	l.Info().Msgf("Adding coin: %s", r.Symbol)

	return m.db.Create(&r)
}
