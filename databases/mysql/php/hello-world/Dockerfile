FROM php

WORKDIR /app

RUN docker-php-ext-install pdo pdo_mysql

COPY version.php .

USER nobody

ENTRYPOINT ["php", "/app/version.php"]