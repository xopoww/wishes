#!/bin/bash

if [ $# -eq 0 ] ; then
    echo "missing required position arg"
    exit 1
fi
FILE=$1

if [ -f $FILE ] ; then
    echo "file already exists"
    exit 1
fi
touch $FILE

migrate -path internal/db/migrations/ -database sqlite3://${FILE} up 