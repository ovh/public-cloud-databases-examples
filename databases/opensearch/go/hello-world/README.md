# Run a hello world on OpenSearch with go

The purpose of this tutorial is to execute a simple command on the opensearch cluster.

## Requirements

You need first to have an OpenSearch cluster deployed. [More information here](https://www.ovhcloud.com/fr/public-cloud/opensearch/)

## Build and run

Once you have created your user, and set up your IP restrictions you can run the following:

```console
read USER
read -s PASSWORD
read URI
docker build -t opensearch-go-hello-world .
docker run opensearch-go-hello-world "$(echo -n $URI | sed "s/<username>/$USER/" | sed "s/<password>/$PASSWORD/")"
```

## Dependencies

This example is using the following external dependencies:
* [github.com/opensearch-project/opensearch-go/v2](https://github.com/opensearch-project/opensearch-go), License Apache-2.0