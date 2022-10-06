variable "ovh" {
  type = map(any)
  default = {
    endpoint           = "ovh-eu"
    application_key    = ""
    application_secret = ""
    consumer_key       = ""
  }
}

variable "product" {
  type = map(any)
  default = {
    project_id = ""
    region     = "DE"
    plan       = "essential"
    flavor     = "db1-7"
    version    = "4.0"
  }
}

variable "access" {
  type = map(any)
  default = {
    ip = "xx.xx.xx.xx/32"
  }
}
