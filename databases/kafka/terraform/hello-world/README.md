# Hello world for Kafka with terraform

The purpose of this tutorial is to create an apache kafka database.

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
    plan       = "business"
    flavor     = "db1-7"
    version    = "3.1"
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

### Delete the cluster

```console
terraform destroy -var-file=secrets.tfvars -auto-approve
```
