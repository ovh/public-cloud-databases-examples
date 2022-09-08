FROM php:7.1.11-apache


WORKDIR /app

ENV CASSANDRA_PHP_DRIVER_VERSION 1.3.2

RUN rm /var/lib/apt/lists/* -vrf && \
    apt-get clean && apt-get update

RUN apt-get update && \
    apt-get install -y git wget libssl-dev zlib1g-dev libicu-dev g++ make cmake libuv-dev libgmp-dev uuid-dev && \
    apt-get autoclean -y && \
    apt-get clean -y

# Install datastax php-driver fox cassandra
RUN git clone https://github.com/datastax/php-driver.git /usr/src/datastax-php-driver && \
    cd /usr/src/datastax-php-driver && \
    git checkout --detach v${CASSANDRA_PHP_DRIVER_VERSION} && \
    git submodule update --init && \
    cd ext && \
    ./install.sh && \
    echo extension=cassandra.so > /usr/local/etc/php/conf.d/cassandra.ini
COPY version.php .
USER nobody

ENTRYPOINT ["php", "/app/version.php"]