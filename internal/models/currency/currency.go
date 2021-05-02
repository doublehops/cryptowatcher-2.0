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

func (m *Model) GetRecordBySymbol(record *database.Currency, s string) {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	m.db.Find(&record, "symbol = ?", s)
}

func (m *Model) GetRecordsMapKeySymbol(curMap *map[string]uint32) {

	var records  []database.Currency
	//mp := make(map[string]uint32)

	l := m.l.Lg.With().Str("currency", "GetRecordIdsAndSymbols").Logger()
	l.Info().Msgf("Fetching currencies attrs of just ID and Symbol")

	m.db.Select("id", "symbol").Find(&records)

	for _, v := range records {
		(*curMap)[v.Symbol] = v.ID
	}
}

func (m *Model) CreateRecord(record *database.Currency) (error) {

	l := m.l.Lg.With().Str("currency", "CreateCurrency").Logger()
	l.Info().Msgf("Adding currency: %s", record.Symbol)

	result := m.db.Create(record)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
		return result.Error
	}

	return nil
}
