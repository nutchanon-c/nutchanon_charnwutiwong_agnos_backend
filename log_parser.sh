#!/bin/bash

# Ensure jq is installed
if ! [ -x "$(command -v jq)" ]; then
  echo 'Error: jq is not installed.' >&2
  exit 1
fi

tail -F /var/log/nginx/access.log | while read line; do
  psql -h "$PGHOST" -U "$PGUSER" -d "$PGDATABASE" -c "INSERT INTO access_logs (request_time, remote_addr, request_method, request_uri, response_status, response_time) VALUES ('$(echo $line | jq -r .request_time)', '$(echo $line | jq -r .remote_addr)', '$(echo $line | jq -r .request_method)', '$(echo $line | jq -r .request_uri)', $(echo $line | jq -r .response_status), $(echo $line | jq -r .response_time));"
done
