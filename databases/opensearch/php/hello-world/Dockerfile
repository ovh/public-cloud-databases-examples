FROM php

COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

WORKDIR /app

RUN apt-get update && apt-get -y install git p7zip \
    && composer require opensearch-project/opensearch-php

COPY hello-world.php .

USER nobody

ENTRYPOINT ["php", "/app/hello-world.php"]