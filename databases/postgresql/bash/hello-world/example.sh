#!/bin/bash

if [[ -z "$1" ]]
then
    echo "URI is missing in the environment" 1>&2
    exit 1
fi

echo "SELECT version();" | psql $1