# Hello world for MySQL with go

The purpose of this tutorial is to execute a simple command on a MySQL database.

## Requirements

You need first to have a MySQL service available. [More information here](https://www.ovhcloud.com/en/public-cloud/mysql/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t mysql-go-hello-world .
docker run mysql-go-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```
