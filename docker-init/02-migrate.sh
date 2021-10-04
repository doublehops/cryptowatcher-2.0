echo "Importing data to MySQL"
mysql -u dev -ppass12 -h 127.0.0.1 cw < database/migrations/000001_init_schema.up.sql
mysql -u dev -ppass12 -h 127.0.0.1 cw_test < database/migrations/000001_init_schema.up.sql