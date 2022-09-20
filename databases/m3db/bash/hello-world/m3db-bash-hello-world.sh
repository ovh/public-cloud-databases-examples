#!/bin/bash

set -e

if [[ -z "$1" ]]
then
    echo "usage: $0 <URI>" 1>&2
    exit 1
fi

curl -s $1/api/v1/labels