#!/bin/bash


psql -d postgres -c "CREATE ROLE phonebook WITH LOGIN PASSWORD 'phonebook';"
psql -d postgres -c "ALTER ROLE phonebook CREATEDB;"
psql -d postgres -c "GRANT pg_read_server_files TO phonebook;"
