FROM mariadb:latest

ENV MYSQL_ALLOW_EMPTY_PASSWORD=1

COPY initdb.d/* /docker-entrypoint-initdb.d/
