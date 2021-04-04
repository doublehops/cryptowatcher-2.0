#!/bin/bash

# Run migrations with Python `migrate` tool. Install with `apt install python3-migrate`.

migrate -source file:./database/migrations -database mysql://localhost:3306/cw $1 $2