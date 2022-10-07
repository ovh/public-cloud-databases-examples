# Hello world for Cassandra with terraform

The purpose of this tutorial is to create a cassandra database.

## Requirements

You need the following:
* [terraform](https://www.terraform.io/) installed
* an [OVHcloud Public cloud project](https://www.ovhcloud.com/en/public-cloud/)
* OVHcloud API credentials
    * [EU](https://www.ovh.com/auth/?onsuccess=https%3A%2F%2Fwww.ovh.com%2Fauth%2FcreateToken%2F%3F)
    * [CA](https://ca.ovh.com/auth/?onsuccess=https%3A//ca.ovh.com%2Fauth%2FcreateToken%2F%3F)

## Build and run

### Create the terraform variables file
```console
read -s APPLICATION_KEY
read -s APPLICATION_SECRET
read -s CONSUMER_KEY
read CLOUD_PROJECT_ID

cat << EOF > secrets.tfvars
ovh = {
    endpoint           = "ovh-eu"
    application_key    = "$APPLICATION_KEY"
    application_secret = "$APPLICATION_SECRET"
    consumer_key       = "$CONSUMER_KEY"
}

product = {
    project_id = "$CLOUD_PROJECT_ID"
    region     = "DE"
    plan       = "essential"
    flavor     = "db1-7"
    version    = "4.0"
}

access = {
    ip  = "$(curl ifconfig.me)/32"
}
EOF
```

### Validate the configuration

```console
terraform init
terraform plan -var-file=secrets.tfvars
```

### Create the cluster 

```console
terraform apply -var-file=secrets.tfvars -auto-approve
```

### Export the credentials

If you need to re-use the credentials in other scripts, you can export the user credentials and the URI

```console
## PATH_TO_CERTIFICATES will be the directory containing all the 3 needed certificates
read PATH_TO_CERTIFICATES
mkdir -p $PATH_TO_CERTIFICATES
terraform output -raw cluster_ca > $PATH_TO_CERTIFICATES/ca.certificate.pem

export PASSWORD=$(terraform output -raw user_password)
export USER=$(terraform output -raw user_name)
export URI=$(terraform output -raw cluster_uri)
```

With these exports you can go directly in any other example (e.g: go) to docker build and run it and see it working.

### Delete the cluster

```console
terraform destroy -var-file=secrets.tfvars -auto-approve
```
