FROM php

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka1 librdkafka-dev
RUN pecl install rdkafka
RUN echo "extension=rdkafka.so" > /usr/local/etc/php/conf.d/rdkafka.ini

COPY kafka-php-hello-world.php .

USER nobody

ENTRYPOINT ["php", "/app/kafka-php-hello-world.php"]