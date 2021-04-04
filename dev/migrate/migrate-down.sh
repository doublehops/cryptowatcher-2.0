#!/bin/bash

migrate -path database/migrations -database "mysql://root:my123@tcp(localhost:3306)/cw" -verbose down