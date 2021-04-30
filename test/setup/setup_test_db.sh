#!/bin/bash

# Prepare test database for tests.

mysql -u root -pmysql -u root -pmy123 -h 127.0.0.1 -e 'DROP DATABASE cw_test; CREATE DATABASE cw_test'
mysql -u root -pmysql -u root -pmy123 -h 127.0.0.1 -e "CREATE USER 'cw_test'@'%' IDENTIFIED BY 'my123'"
mysql -u root -pmysql -u root -pmy123 -h 127.0.0.1 -e "FLUSH PRIVILEGES"
mysql -u root -pmysql -u root -pmy123 -h 127.0.0.1 -e "GRANT ALL PRIVILEGES ON cw_test.* TO 'cw_test'@'%'"
#
migrate -path database/migrations -database "mysql://cw_test:my123@tcp(localhost:3306)/cw_test" -verbose up
