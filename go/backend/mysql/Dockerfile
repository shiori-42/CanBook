FROM mysql:8.0

COPY docker-entrypoint-initdb.d /docker-entrypoint-initdb.d

ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE db

EXPOSE 3306