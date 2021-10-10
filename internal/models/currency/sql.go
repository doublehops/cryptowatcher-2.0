package currency

var GetRecordsSql = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  ORDER BY ID
  LIMIT ?,?
`

var GetRecordByIDSql = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE id = ?`

var GetRecordBySymbolSql = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE symbol = ?`

var InsertRecordSql = `
INSERT INTO currency
(name, symbol, created_at, updated_at)
VALUES
(?, ?, NOW(), NOW())
`
