#!/bin/bash

set -e

if [ $# -ne 1 ]
  then
    echo "URI is missing"
    exit 1
fi

URI=$(echo "$1" | sed -r 's/^(.*:\/\/)(.*)/\2/g')
USERPASS=$(echo "$URI" | cut -d"@" -f1)
USER=$(echo "$USERPASS" | cut -d":" -f1)
PASS=$(echo "$USERPASS" | cut -d":" -f2)
ENDPOINT=$(echo "${URI/$USERPASS@/}" | cut -d"/" -f1)

configFile=$(mktemp)

cat << EOF > $configFile
profiles:
  - name: ovhcloud
    endpoint: 'https://$ENDPOINT'
    user: '$USER'
    password: '$PASS'
    max_retry: 3
    timeout: 10
EOF

opensearch-cli --config $configFile --profile ovhcloud curl get --path _nodes | jq -r '.nodes[] | "\(.name): \(.version)"'