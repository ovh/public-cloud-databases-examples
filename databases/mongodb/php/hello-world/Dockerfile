FROM php

WORKDIR /app

RUN apt-get update && apt-get install -y openssl libssl-dev libcurl4-openssl-dev
RUN pecl install mongodb
RUN echo "extension=mongodb.so" > /usr/local/etc/php/conf.d/mongo.ini

COPY version.php .

USER nobody

ENTRYPOINT ["php", "/app/version.php"]