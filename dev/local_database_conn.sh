#!/bin/bash

# Connect to dev database.

source .env
mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE