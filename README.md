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
UP: dev/migrate/migrate.sh
Down: dev/migrate/migrate-down.sh
```