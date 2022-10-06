output "cluster_ca" {
  value = data.ovh_cloud_project_database_certificates.certificates.ca
}

output "user_password" {
  value     = ovh_cloud_project_database_user.userjd.password
  sensitive = true
}

output "user_name" {
  value = ovh_cloud_project_database_user.userjd.name
}

output "cluster_uri" {
  value = ovh_cloud_project_database.service.endpoints.0.uri
}


