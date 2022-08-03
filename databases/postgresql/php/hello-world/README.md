# Run a hello world on Postgresql with php

The purpose of this tutorial is to execute a simple command on the postgresql cluster.

## Requirements

You need first to have a Postgresql cluster deployed. [More information here](https://www.ovhcloud.com/fr/public-cloud/postgresql/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t postgresql-php-hello-world .
docker run postgresql-php-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```