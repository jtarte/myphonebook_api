FROM alpine

RUN apk update && apk add "postgresql-client" && apk add "bash"

ADD load.sh .
ADD mygroup.csv .
ADD person.csv .



#ENTRYPOINT [ "psql", "-h", "$DB_HOST", "-d", "$DB_NAME", "-U", "$DB_USER"]
ENTRYPOINT ["./load.sh"]