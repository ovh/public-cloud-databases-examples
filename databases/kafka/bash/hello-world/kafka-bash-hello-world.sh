#!/bin/bash

set -e

if [[ -z "$1" ]]
then
    echo "usage: $0 <URI>" 1>&2
    exit 1
fi

configFile=$(mktemp)

cat << EOF > $configFile
bootstrap.servers=$1

security.protocol=ssl

ssl.ca.location=/certificates/ca.certificate.pem
ssl.certificate.location=/certificates/access.certificate.pem
ssl.key.location=/certificates/access.key
EOF

kafkacat -F $configFile -L