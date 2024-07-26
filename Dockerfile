FROM mysql

ENV MYSQL_ROOT_PASSWORD=qazwsx 
ENV MYSQL_DATABASE=jwt_gin
ENV MYSQL_USER=thiago
ENV MYSQL_PASSWORD=qazwsx

COPY ./db/ /docker-entrypoint-initdb.d/

EXPOSE 3306
