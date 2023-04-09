FROM ubuntu:latest

RUN apt-get update \
  && apt-get install -y mysql-server \
  && apt-get install -y systemd \
  && apt-get install -y vim

EXPOSE 3306

CMD ["mysqld_safe"]
