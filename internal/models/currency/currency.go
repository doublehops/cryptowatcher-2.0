package currency

import (
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"gorm.io/gorm"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db *gorm.DB
	l *logga.Logga
}

// New - creates new instance of currency.
func New(db *gorm.DB, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l: logger,
	}
}

// GetRecordBySymbol will return the requested record from the database by its ID.
func (m *Model) GetRecordBySymbol(record *database.Currency, s string) {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	m.db.Find(&record, "symbol = ?", s)
}

// GetRecords will return model records.
func (m *Model) GetRecords(records *database.Currencies, pg *pagination.Meta) {

	l := m.l.Lg.With().Str("currency", "GetRecords").Logger()
	l.Info().Msgf("Fetching currencies")

	m.db.Debug().Limit(pg.PerPage).Offset(pg.Offset).Find(records)
}

// GetRecordsMapKeySymbol will return the requested record from the database by its symbol.
func (m *Model) GetRecordsMapKeySymbol(curMap *map[string]uint32) {

	var records  []database.Currency

	l := m.l.Lg.With().Str("currency", "GetRecordIdsAndSymbols").Logger()
	l.Info().Msgf("Fetching currencies attrs of just ID and Symbol")

	m.db.Debug().Select("id", "symbol").Find(&records)

	for _, v := range records {
		(*curMap)[v.Symbol] = v.ID
	}
}

// CreateRecord will create a new record in the database.
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
