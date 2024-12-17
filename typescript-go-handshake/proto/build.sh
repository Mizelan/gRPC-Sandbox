#!/bin/sh
set -e
set -o pipefail

buf generate --output ../gen
buf build -o ../gen/user.pb

mkdir -p ../go-server/gen && cp -r ../gen/go/* ../go-server/gen/
mkdir -p ../node-client/src/gen && cp -r ../gen/ts/* ../node-client/src/gen/
mkdir -p ../nextjs-backend-relay/src/gen && cp -r ../gen/ts/* ../nextjs-backend-relay/src/gen/