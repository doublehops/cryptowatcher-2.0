package cmchistory

// TimeSeriesSlicedPeriodQuery - Get average quote_price for a given period
var TimeSeriesSlicedPeriodQuery = `
	SELECT 
		AVG(c.quote_price) AS quote_price,
		c.created_at
	FROM
		(SELECT *, NTILE(?) OVER (ORDER BY created_at) AS bucket FROM cmc_history) AS c
	WHERE c.symbol = ? 
	AND c.created_at >= ? 
	AND c.created_at <= ?
	GROUP BY bucket
`

var GetRecordsSql = `
	SELECT id,symbol,name,created_at,updated_at FROM cmc_history
	ORDER BY ID
	LIMIT ?,?
`

var GetRecordByIDSql = `
SELECT 
	id,
    currency_id,
	name,
	symbol,
	slug,
	num_market_pairs,
	date_added,
	max_supply,
	circulating_supply,
	total_supply,
	cmc_rank,
	quote_price,
	volume_24h,
	percent_change_1h,
	percent_change_24h,
	percent_change_7d,
	percent_change_30d,
	percent_change_60d,
	percent_change_90d,
	market_cap,
	created_at,
	updated_at
  FROM cmc_history
  WHERE id = ?`

var GetRecordBySymbolSql = `
SELECT 
	id,
    currency_id,
	name,
	symbol,
	slug,
	num_market_pairs,
	date_added,
	max_supply,
	circulating_supply,
	total_supply,
	cmc_rank,
	quote_price,
	volume_24h,
	percent_change_1h,
	percent_change_24h,
	percent_change_7d,
	percent_change_30d,
	percent_change_60d,
	percent_change_90d,
	market_cap,
	created_at,
	updated_at
  FROM cmc_history AS 
  WHERE symbol = ?`

var InsertRecordSql = `
INSERT INTO cmc_history (
	currency_id,
	name,
	symbol,
	slug,
	num_market_pairs,
	date_added,
	max_supply,
	circulating_supply,
	total_supply,
	cmc_rank,
	quote_price,
	volume_24h,
	percent_change_1h,
	percent_change_24h,
	percent_change_7d,
	percent_change_30d,
	percent_change_60d,
	percent_change_90d,
	market_cap,
	created_at,
	updated_at
)
VALUES (
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	?,
	NOW(),
	NOW()
)
`
