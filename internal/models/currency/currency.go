package currency

import (
	dbi "cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/pkg/handlers/pagination"
	"database/sql"
	"fmt"
	"reflect"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db dbi.QueryAble
	l  *logga.Logga
}

// New - creates new instance of currency.
func New(db dbi.QueryAble, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l:  logger,
	}
}

// GetRecordByID will return the requested record from the db by its ID.
func (m *Model) GetRecordByID(record *database.Currency, ID int64) error {

	l := m.l.Lg.With().Str("currency", "GetCoinByID").Logger()
	l.Info().Msgf("Fetching currency by ID: %d", ID)

	row := m.db.QueryRow(GetRecordByIDSql, ID)
	err := m.populateRecord(record, row)
	if err != nil {
		return fmt.Errorf("unable to populate record. %s", err)
	}

	return nil
}

// GetRecordBySymbol will return a record by its symbol.
func (m *Model) GetRecordBySymbol(record *database.Currency, s string) error {

	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	row := m.db.QueryRow(GetRecordBySymbolSql, s)
	err := m.populateRecord(record, row)
	if err != nil {
		return fmt.Errorf("unable to populate record. %s", err)
	}

	return nil
}

// GetRecords will return model records.
func (m *Model) GetRecords(pg *pagination.MetaRequest) (database.Currencies, error) {

	l := m.l.Lg.With().Str("currency", "GetRecords").Logger()
	l.Info().Msgf("Fetching currencies")

	var records database.Currencies
	rows, err := m.db.Query(GetRecordsSql, pg.Offset, pg.PerPage)
	if err != nil {
		err := fmt.Errorf("unable to retrieve currency records. %w", err)
		l.Error().Msg(err.Error())
		return records, err
	}
	defer rows.Close()

	for rows.Next() {
		var record database.Currency
		if err := rows.Scan(&record.ID, &record.Symbol, &record.Name, &record.CreatedAt, &record.UpdatedAt); err != nil {
			return records, fmt.Errorf("error scanning row. %w", err)
		}

		records = append(records, &record)
	}

	return records, nil
}

// GetRecordsMapKeySymbol will return the requested record from the db by its symbol.
func (m *Model) GetRecordsMapKeySymbol() (map[string]uint32, error) {

	l := m.l.Lg.With().Str("currency", "GetRecordsMapKeySymbol").Logger()
	l.Info().Msgf("Fetching currencies attrs of just ID and Symbol")

	curMap := make(map[string]uint32)
	pg := pagination.MetaRequest{
		Page: 1,
		PerPage: 100000,
		Offset: 0,
	}

	records, err := m.GetRecords(&pg)
	if err != nil {
		return curMap, err
	}

	for _, v := range records {
		curMap[v.Symbol] = v.ID
	}

	return curMap, nil
}

// CreateRecord will create a new record in the db.
func (m *Model) CreateRecord(record *database.Currency) (int64, error) {

	l := m.l.Lg.With().Str("currency", "CreateCurrency").Logger()
	l.Info().Msgf("Adding currency: %s; with interface type: %v", record.Symbol, reflect.TypeOf(m.db))

	result, err := m.db.Exec(InsertRecordSql, record.Name, record.Symbol)
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
