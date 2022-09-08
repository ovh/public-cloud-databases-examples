#!/bin/bash

set -e


if [ $# -ne 4 ]
  then
    echo "Must give 4 arguments: Host, Port, Username, Password"
    exit 1
fi

export CQLSH_PORT=$2


echo "SELECT cql_version FROM system.local; exit" | cqlsh $1 -u $3 -p $4 --ssl
