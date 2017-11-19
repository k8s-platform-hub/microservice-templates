FROM java:8

RUN apt-get update

RUN apt-get -y install nginx

RUN wget https://github.com/schemaspy/schemaspy/releases/download/v6.0.0-rc2/schemaspy-6.0.0-rc2.jar -O schemaSpy.jar

RUN wget https://jdbc.postgresql.org/download/postgresql-42.1.4.jar -O postgresql-jdbc4.jar

RUN apt-get install -y graphviz

RUN rm -v /etc/nginx/nginx.conf

RUN useradd --no-create-home nginx

COPY app/conf/nginx.conf /etc/nginx

COPY app/init.sh /bin/

RUN chmod a+x /bin/init.sh

CMD ["/bin/init.sh"]
