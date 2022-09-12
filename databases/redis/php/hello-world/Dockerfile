FROM php

WORKDIR /app

RUN pecl install redis && docker-php-ext-enable redis

COPY version.php .

USER nobody

ENTRYPOINT ["php", "/app/version.php"]
