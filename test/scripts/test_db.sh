#!/bin/bash

# Start the PostgreSQL container
docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=585858 -p 6432:5432 -d postgres:latest
echo "PostgreSQL starting..."
sleep 5  # Wait for the container to fully start

# Create the database
docker exec postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
if [ $? -ne 0 ]; then
    echo "Failed to create database"
    exit 1
fi
echo "Database productapp created"

# Create the table
docker exec postgres-test psql -U postgres -d productapp -c "
CREATE TABLE IF NOT EXISTS products (
  id bigserial NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL,
  price double precision NOT NULL,
  discount double precision,
  store varchar(255) NOT NULL
);"

if [ $? -ne 0 ]; then
    echo "Failed to create table"
    exit 1
fi

echo "Table products created"
