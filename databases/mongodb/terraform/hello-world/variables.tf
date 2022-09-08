variable "ovh" {
  type = map(string)
  default = {
    endpoint           = "ovh-eu"
    application_key    = ""
    application_secret = ""
    consumer_key       = ""
  }
}

variable "product" {
  type = map(string)
  default = {
    project_id = ""
    region     = "DE"
    plan       = "essential"
    flavor     = "db1-7"
    version    = "5.0"
  }
}

variable "access" {
  type = map(string)
  default = {
    ip = "xx.xx.xx.xx/32"
  }
}
