package currency

var GetRecordsSQL = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  ORDER BY ID
  LIMIT ?,?
`

var GetRecordByIDSQL = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE id = ?`

var GetRecordBySymbolSQL = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE symbol = ?`

var InsertRecordSQL = `
INSERT INTO currency
(name, symbol, created_at, updated_at)
VALUES
(?, ?, NOW(), NOW())
`

var DeleteRecordSQL = `
DELETE FROM currency
WHERE id=?
`
