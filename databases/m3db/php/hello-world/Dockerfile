FROM php

COPY --from=composer /usr/bin/composer /usr/bin/

WORKDIR /app

RUN apt-get update && apt-get install -y git
RUN composer require guzzlehttp/guzzle

COPY m3db-php-hello-world.php .

USER nobody

ENTRYPOINT ["php", "/app/m3db-php-hello-world.php"]