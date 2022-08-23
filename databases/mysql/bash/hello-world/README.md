# Hello world for MySQL with bash

The purpose of this tutorial is to execute a simple command on a MySQL database.

## Requirements

You need first to have a MySQL service available. [More information here](https://www.ovhcloud.com/en/public-cloud/mysql/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t mysql-bash-hello-world .
docker run mysql-bash-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```
