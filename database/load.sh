#!/bin/bash

export PGPASSWORD=$DB_PWD

psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_USER -c "CREATE TABLE mygroup(id SERIAL PRIMARY KEY, name VARCHAR(30));"
psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_USER -c "CREATE TABLE person(id SERIAL PRIMARY KEY, firstname VARCHAR(30), name VARCHAR(30),email VARCHAR(30), phonenumber varchar(15), gr_id INTEGER REFERENCES mygroup(id));"
psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_NAME -c "\copy mygroup (name) from './mygroup.csv' DELIMITER ',' CSV HEADER;"
psql -h $DB_HOST -p $DB_PORT -d $DB_NAME -U $DB_NAME -c "\copy person (firstname,name,email,phonenumber,gr_id) from './person.csv' DELIMITER ',' CSV HEADER;"