#! /bin/bash

set -e

# Stop and remove old containers
docker stop sqlboilercontainer || true
docker rm sqlboilercontainer || true

# Compile source
go build -o app

# Build and run docker image
docker build --tag sqlboiler:replicate-issue --no-cache --file Dockerfile-db .
docker run --name sqlboilercontainer --detach --publish 5432:5432 sqlboiler:replicate-issue

# Give time for detached container to fully start-up
sleep 3

# Generate the models
sqlboiler --wipe psql

# Run the application
./app
