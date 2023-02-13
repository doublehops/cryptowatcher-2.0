### Get Running
Run `make docker_up` to get the local environment up.

### Database Migrations

Run `make migrate` to run migrations.

#### Run tests

Testing records go into a test database `cw_test`. Before testing, the database should be
setup with the command `test/setup/setup_test_db.sh`. Then tests can be run with:

```
go test ./...
```

#### ADDITIONAL THINGS
cURL request to get history from CMC.
```
curl -H "X-CMC_PRO_API_KEY: _YOUR_KEY_" "Accept: application/json" -d "start=1&limit=5000&convert=USD" -G https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
```

Debug in GORM:
```
result := m.db.Debug().Create(&record)
```
