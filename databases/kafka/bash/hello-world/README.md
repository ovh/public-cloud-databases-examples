# Run a hello world on Kafka with bash

The purpose of this tutorial is to execute a simple command on the apache kafka cluster.

## Requirements

You need first to have a apache kafka cluster deployed. [More information here](https://www.ovhcloud.com/fr/public-cloud/apache-kafka/)

### Certificates
You need to create a directory that will be mounted on the docker container. It must contains the following

*  `ca.certificate.pem`: the cluster certificate
    * GET [/cloud/project/{serviceName}/database/kafka/{clusterId}/certificates](https://api.ovh.com/console/#/cloud/project/{serviceName}/database/kafka/{clusterId}/certificates~GET) or
    * Available on the landing page of your cluster in the manager

*  `access.certificate.pem`: the user certificate
    *  GET [/cloud/project/{serviceName}/database/kafka/{clusterId}/user/{userId}/access](https://api.ovh.com/console/#/cloud/project/{serviceName}/database/kafka/{clusterId}/user/{userId}/access~GET) or
    * Click `Show the certificate` in the manager in the User section

*  `access.key`: the user key
    *  GET [/cloud/project/{serviceName}/database/kafka/{clusterId}/user/{userId}/access](https://api.ovh.com/console/#/cloud/project/{serviceName}/database/kafka/{clusterId}/user/{userId}/access~GET)
    * Click `Show the certificate` in the manager in the User section

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read URI
## PATH_TO_CERTIFICATES must be a directory containing all the 3 certificates listed previously
read PATH_TO_CERTIFICATES
docker build -t kafka-bash-hello-world .
docker run -v $PATH_TO_CERTIFICATES:/certificates kafka-bash-hello-world $URI
```

## Dependencies

This example is using the following external dependencies:
* [confluentinc/cp-kafkacat/](https://hub.docker.com/r/confluentinc/cp-kafkacat/), License Apache-2.0