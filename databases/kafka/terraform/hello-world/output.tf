output "cluster_ca" {
  value = data.ovh_cloud_project_database_certificates.certificates.ca
}

output "cluster_uri" {
  value = ovh_cloud_project_database.service.endpoints.0.uri
}

output "user_cert" {
  value = data.ovh_cloud_project_database_kafka_user_access.access.cert
}

output "user_key" {
  value     = data.ovh_cloud_project_database_kafka_user_access.access.key
  sensitive = true
}

output "user_name" {
  value = ovh_cloud_project_database_user.userjd.name
}

output "user_password" {
  value     = ovh_cloud_project_database_user.userjd.password
  sensitive = true
}


