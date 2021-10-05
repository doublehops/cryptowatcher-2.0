package currency

import (
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"database/sql"
	"fmt"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db *sql.DB
	l *logga.Logga
}

// New - creates new instance of currency.
func New(db *sql.DB, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l: logger,
	}
}

// GetRecordByID will return the requested record from the db by its ID.
func (m *Model) GetRecordByID(record *database.Currency, ID int64) error {

	l := m.l.Lg.With().Str("currency", "GetCoinByID").Logger()
	l.Info().Msgf("Fetching currency by ID: %d", ID)

	//bindVars := map[string]interface{}{
	//	"symbol": s,
	//}
	//row := m.db.QueryRow(GetRecordBySymbol, s)
	//err := row.Scan(record, s)

	//err := db.QueryToStructs(record, m.db, GetRecordByID, ID)
	row := m.db.QueryRow(GetRecordByID, ID)
	err := m.populateRecord(record, row)
	if err != nil {
		return fmt.Errorf("unable to populate record. %s", err)
	}

	return nil
}

// GetRecordBySymbol will return the requested record from the db by its symbol.
func (m *Model) GetRecordBySymbol(record *database.Currency, s string) error {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	//bindVars := map[string]interface{}{
	//	"symbol": s,
	//}
	row := m.db.QueryRow(GetRecordBySymbol, s)
	err := m.populateRecord(record, row)
	if err != nil {
		return fmt.Errorf("unable to populate record. %s", err)
	}

	return nil
}

// GetRecords will return model records.
func (m *Model) GetRecords(records *database.Currencies, pg *pagination.MetaRequest, count *int64) {

	l := m.l.Lg.With().Str("currency", "GetRecords").Logger()
	l.Info().Msgf("Fetching currencies")

	//m.db.Find(records).Count(count)
	//m.db.Limit(pg.PerPage).Offset(pg.Offset).Find(records)
}

// GetRecordsMapKeySymbol will return the requested record from the db by its symbol.
func (m *Model) GetRecordsMapKeySymbol(curMap *map[string]uint32) {

	var records  []database.Currency

	l := m.l.Lg.With().Str("currency", "GetRecordIdsAndSymbols").Logger()
	l.Info().Msgf("Fetching currencies attrs of just ID and Symbol")

	for _, v := range records {
		(*curMap)[v.Symbol] = v.ID
	}
}

// CreateRecord will create a new record in the db.
func (m *Model) CreateRecord(record *database.Currency) (int64, error) {

	l := m.l.Lg.With().Str("currency", "CreateCurrency").Logger()
	l.Info().Msgf("Adding currency: %s", record.Symbol)

	result, err := m.db.Exec(InsertRecord, record.Name, record.Symbol)
	if err != nil {
		l.Error().Msgf("There was an error saving record to db. %w", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// populateRecord will populate model object from query.
func (m *Model) populateRecord(record *database.Currency, row *sql.Row) error {

	err := row.Scan(&record.ID, &record.Symbol, &record.Name, &record.CreatedAt, &record.UpdatedAt)

	return err
}
