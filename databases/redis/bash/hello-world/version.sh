#!/bin/bash

set -e

if [[ -z "$1" ]]
then
    echo "URI is missing" 1>&2
    exit 1
fi

version=$(redis-cli -u $1 info | grep redis_version)
echo $version
