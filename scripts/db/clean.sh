#!/bin/bash

if [ $# -eq 0 ] ; then
    echo "missing required position arg"
    exit 1
fi
FILE=$1

migrate -path internal/db/migrations/ -database sqlite3://${FILE} down
migrate -path internal/db/migrations/ -database sqlite3://${FILE} up