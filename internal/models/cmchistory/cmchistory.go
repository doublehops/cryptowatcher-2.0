package cmchistory

import (
	"database/sql"

	dbi "cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db dbi.QueryAble
	l  *logga.Logga
}

type SearchParams struct {
	TimeFrom     string
	TimeTo       string
	TimeFromUnix int64
	TimeToUnix   int64
}

// New - creates new instance of cmchistory.
func New(db dbi.QueryAble, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l:  logger,
	}
}

// CreateRecord inserts a new record into the cmc_history table.
func (m *Model) CreateRecord(record *database.CmcHistory) (uint32, error) {

	l := m.l.Lg.With().Str("cmchistory", "CreateRecord").Logger()
	l.Info().Msgf("Adding cmc record: %s", record.Symbol)

	result, err := m.db.Exec(InsertRecordSql, &record.CurrencyID,
		&record.Name,
		&record.Symbol,
		&record.Slug,
		&record.NumMarketPairs,
		&record.DateAdded,
		&record.MaxSupply,
		&record.CirculatingSupply,
		&record.TotalSupply,
		&record.CmcRank,
		&record.QuotePrice,
		&record.Volume24h,
		&record.PercentChange1h,
		&record.PercentChange24h,
		&record.PercentChange7D,
		&record.PercentChange30D,
		&record.PercentChange60D,
		&record.PercentChange90D,
		&record.MarketCap,
	)
	if err != nil {
		l.Error().Msgf("There was an error saving record to db. %w", err)
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	l.Info().Msgf("Added cmc record. Symbol: %s; ID: %d", record.Symbol, lastInsertID)

	return uint32(lastInsertID), err
}

// GetRecordByID will return the requested record from the db by its ID.
func (m *Model) GetRecordByID(record *database.CmcHistory, ID uint32) error {

	l := m.l.Lg.With().Str("cmchistory", "GetRecordByID").Logger()
	l.Info().Msgf("Retrieving cmchistory record by ID: %d", ID)

	row := m.db.QueryRow(GetRecordByIDSql, ID)
	err := m.populateRecord(record, row)
	if err != nil {
		l.Error().Msgf("There was an error retrieving record from the db. %w", err)
		return err
	}

	return nil
}

// GetPriceTimeSeriesData will return records grouped together in X number of groups with `quote_price` averaged out per group/bucket.
func (m *Model) GetPriceTimeSeriesData(symbol string, searchParams *SearchParams) ([]*database.CmcHistoryPriceTimeSeriesDataItem, error) {

	l := m.l.Lg.With().Str("cmchistory", "GetTimeSeriesData").Logger()
	l.Info().Msgf("Fetching cmchistory records for symbol: %s", symbol)

	buckets := 5 // number of buckets to group the records in and determine average for.

	rows, err := m.db.Query(TimeSeriesSlicedPeriodQuery, buckets, symbol, searchParams.TimeFrom, searchParams.TimeTo)
	if err != nil {
		l.Error().Msgf("There was an error retrieving time series data. %w", err)
	}
	defer rows.Close()

	var records []*database.CmcHistoryPriceTimeSeriesDataItem
	for rows.Next() {
		var record database.CmcHistoryPriceTimeSeriesDataItem
		err = rows.Scan(&record.QuotePrice, &record.CreatedAt)
		if err != nil {
			return records, err
		}

		records = append(records, &record)
	}

	return records, nil
}

// populateRecord will populate model object from query.
func (m *Model) populateRecord(record *database.CmcHistory, row *sql.Row) error {

	err := row.Scan(&record.ID,
		&record.CurrencyID,
		&record.Name,
		&record.Symbol,
		&record.Slug,
		&record.NumMarketPairs,
		&record.DateAdded,
		&record.MaxSupply,
		&record.CirculatingSupply,
		&record.TotalSupply,
		&record.CmcRank,
		&record.QuotePrice,
		&record.Volume24h,
		&record.PercentChange1h,
		&record.PercentChange24h,
		&record.PercentChange7D,
		&record.PercentChange30D,
		&record.PercentChange60D,
		&record.PercentChange90D,
		&record.MarketCap,
		&record.CreatedAt,
		&record.UpdatedAt)

	return err
}
