# Hello world for M3DB with bash

The purpose of this tutorial is to execute a simple command on a M3DB database.

## Requirements

You need first to have a m3db service available. [More information here](https://www.ovhcloud.com/en/public-cloud/m3db/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t m3db-bash-hello-world .
docker run m3db-bash-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```
