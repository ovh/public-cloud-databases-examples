#!/bin/bash

set -e

if [[ -z "$1" ]]
then
    echo "URI is missing" 1>&2
    exit 1
fi

mysqlsh --uri $1 --sql --table --execute 'select version();'
