# Hello world for Cassandra with php

The purpose of this tutorial is to execute a simple command on a Cassandra database.

## Requirements

You need first to have a Cassandra service available. [More information here](https://www.ovhcloud.com/en/public-cloud/apache-cassandra/)

### Certificate
You need to create a directory that will be mounted on the docker container. It must contains the following:
*  `ca.certificate.pem`: the cluster certificate
    * GET [/cloud/project/{serviceName}/database/cassandra/{clusterId}/certificates](https://api.ovh.com/console/#/cloud/project/{serviceName}/database/cassandra/{clusterId}/certificates~GET) or
    * Available on the landing page of your cluster in the manager

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read HOST
read PORT
## PATH_TO_CERTIFICATES must be a directory containing the certificate
read PATH_TO_CERTIFICATES
docker build -t cassandra-php-hello-world .
docker run -v $PATH_TO_CERTIFICATES:/certificates cassandra-php-hello-world $HOST $PORT $USER $PASSWORD
```

## Dependencies

This example is using the following external dependencies:
* [github.com/datastax/php-driver](https://github.com/datastax/php-driver), License Apache-2.0