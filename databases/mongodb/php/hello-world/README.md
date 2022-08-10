# Hello world for MongoDB with PHP

The purpose of this tutorial is to execute a simple command on a MongoDB database.

## Requirements

You need first to have a MongoDB service available. [More information here](https://www.ovhcloud.com/en/public-cloud/mongodb/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t mongodb-php-hello-world .
docker run mongodb-php-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```
