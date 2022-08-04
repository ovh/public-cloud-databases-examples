FROM php

WORKDIR /app

RUN apt-get update && apt-get install -y libpq-dev && docker-php-ext-install pdo_pgsql

COPY example.php .

USER nobody

ENTRYPOINT ["php", "/app/example.php"]