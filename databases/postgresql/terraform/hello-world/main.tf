terraform {
  required_providers {
    ovh = {
      source  = "ovh/ovh"
      version = "0.20"
    }
  }

  required_version = ">= 0.17.1"
}

provider "ovh" {
  endpoint           = var.ovh.endpoint
  application_key    = var.ovh.application_key
  application_secret = var.ovh.application_secret
  consumer_key       = var.ovh.consumer_key
}

resource "ovh_cloud_project_database" "service" {
  service_name = var.product.project_id
  description  = "terraform-hello-world"
  engine       = "postgresql"
  version      = var.product.version
  plan         = var.product.plan
  nodes {
    region = var.product.region
  }
  flavor = var.product.flavor
}


resource "ovh_cloud_project_database_postgresql_user" "userjd" {
  service_name = ovh_cloud_project_database.service.service_name
  cluster_id   = ovh_cloud_project_database.service.id
  name         = "johndoe"
}

resource "ovh_cloud_project_database_ip_restriction" "iprestriction" {
  service_name = ovh_cloud_project_database.service.service_name
  engine       = ovh_cloud_project_database.service.engine
  cluster_id   = ovh_cloud_project_database.service.id
  ip           = var.access.ip
}
