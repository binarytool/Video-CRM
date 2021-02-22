FROM mysql
ENV MYSQL_ROOT_PASSWORD=1
ENV MYSQL_DATABASE=video-crm

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306

#Just for remember docker run --name phpmyadmin -d --link mysql-server:db -p 8081:80 phpmyadmin/phpmyadmin