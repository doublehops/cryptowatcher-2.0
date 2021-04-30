### Database Migrations

Migrations are run with the migration tool found <a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">here</a>.

#### Install database migration tool

```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update && apt install -y migrate
```

That will likely place `migrate` in `/usr/local/bin/migrate`.

#### Create migrations

```
migrate create -ext sql -dir database/migrations -seq {migration_name}
```

#### Run migrations

```
Up: dev/migrate/migrate.sh
Down: dev/migrate/migrate-down.sh
```

#### Run tests

Tests are run with the following command:

```
go run ./...
```

However, first you may want to run migrations into the test database:
```
migrate -path database/migrations -database "mysql://root:my123@tcp(localhost:3306)/cw_test" -verbose up
```

cURL request to get history from CMC.
```
curl -H "X-CMC_PRO_API_KEY: _YOUR_KEY_" "Accept: application/json" -d "start=1&limit=5000&convert=USD" -G https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
```
