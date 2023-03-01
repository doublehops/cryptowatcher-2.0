package currency

import (
	"database/sql"
	"fmt"
	"reflect"

	dbi "github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/handlers/pagination"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
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

	row := m.db.QueryRow(GetRecordByIDSQL, ID)
	err := m.populateRecord(record, row)
	if err != nil {
		return fmt.Errorf("unable to populate record. %s", err)
	}

	return nil
}

// GetRecordBySymbol will return a record from the database by its symbol.
func (m *Model) GetRecordBySymbol(record *database.Currency, s string) error {
	l := m.l.Lg.With().Str("currency", "GetCoinBySymbol").Logger()
	l.Info().Msgf("Fetching currency by symbol: %s", s)

	row := m.db.QueryRow(GetRecordBySymbolSQL, s)
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
	rows, err := m.db.Query(GetRecordsSQL, pg.Offset, pg.PerPage)
	if err != nil {
		err := fmt.Errorf("unable to retrieve currency records from database. %w", err)
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
		Page:    1,
		PerPage: 100000,
		Offset:  0,
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

	result, err := m.db.Exec(InsertRecordSQL, record.Name, record.Symbol)
	if err != nil {
		l.Error().Msgf("There was an error saving record to db. %s", err)

		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// DeleteRecord will remove a record from the db.
func (m *Model) DeleteRecord(ID uint32) error {
	l := m.l.Lg.With().Str("currency", "DeleteRecord").Logger()
	l.Info().Msgf("Deleting currency with ID %d", ID)

	_, err := m.db.Exec(DeleteRecordSQL, ID)
	if err != nil {
		l.Error().Msgf("There was an error deleting record from db. %s", err)

		return err
	}

	return nil
}

// populateRecord will populate model object from query.
func (m *Model) populateRecord(record *database.Currency, row *sql.Row) error {
	err := row.Scan(&record.ID, &record.Symbol, &record.Name, &record.CreatedAt, &record.UpdatedAt)

	return err
}
