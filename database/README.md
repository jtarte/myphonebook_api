### Database scripts

This deirectory contains script to define and populate the databse used by the API service.

The `load.sh`script is used to create the database schema and load initial data (CSV files).

The `Dockerfile` could be used to generate the container image that could be used for an init job loading the target database. 
