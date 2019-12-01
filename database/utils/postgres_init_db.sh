#!/bin/bash

export PGPASSWORD=$DB_PWD
#psql -d postgres -U phonebook -c "CREATE DATABASE phonebook;"
psql -h $DB_HOST -d $DB_NAME -U $DB_USER -c "CREATE TABLE mygroup(id SERIAL PRIMARY KEY, name VARCHAR(30));"
psql -h $DB_HOST -d $DB_NAME -U $DB_USER-c "CREATE TABLE person(id SERIAL PRIMARY KEY, firstname VARCHAR(30), name VARCHAR(30),email VARCHAR(30), phonenumber varchar(15), gr_id INTEGER REFERENCES mygroup(id));"
