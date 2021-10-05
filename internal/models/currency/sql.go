package currency

var GetRecordByID = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE id = ?`

var GetRecordBySymbol = `
SELECT id,symbol,name,created_at,updated_at FROM currency
  WHERE symbol = ?`

var InsertRecord = `
INSERT INTO currency
(name, symbol, created_at, updated_at)
VALUES
(?, ?, NOW(), NOW())
`
