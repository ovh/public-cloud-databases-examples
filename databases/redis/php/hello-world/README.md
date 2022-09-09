# Hello world for Redis™ with PHP

The purpose of this tutorial is to execute a simple command on a Redis™ database.

## Requirements

You need first to have a Redis™ service available. [More information here](https://www.ovhcloud.com/en/public-cloud/redis/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t redis-php-hello-world .
docker run redis-php-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```
