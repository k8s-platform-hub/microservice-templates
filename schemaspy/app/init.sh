#!/bin/sh
java -jar schemaSpy.jar -t pgsql -s public -db hasuradb -u admin -p $POSTGRES_PASSWORD -host postgres.hasura -o /usr/share/nginx/html -dp postgresql-jdbc4.jar
nginx -g "daemon off;"