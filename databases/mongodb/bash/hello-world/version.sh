#!/bin/bash

set -e

if [[ -z "$1" ]]
then
    echo "URI is missing" 1>&2
    exit 1
fi

mongo --quiet $1 --eval 'db.version()'
