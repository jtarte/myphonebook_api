#!/bin/bash


psql -d postgres -c "DROP DATABASE phonebook;"
psql -d postgres -c "DROP OWNED BY phonebook;"
psql -d postgres -c "DROP ROLE phonebook;"
