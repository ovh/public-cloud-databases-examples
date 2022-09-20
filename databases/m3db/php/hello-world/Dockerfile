FROM php

COPY --from=composer /usr/bin/composer /usr/bin/

WORKDIR /app

RUN apt-get update && apt-get install -y git tree
RUN composer require influxdata/influxdb-client-php nyholm/psr7 php-http/curl-client

RUN tree .

COPY m3db-php-hello-world.php .

USER nobody

ENTRYPOINT ["php", "/app/m3db-php-hello-world.php"]