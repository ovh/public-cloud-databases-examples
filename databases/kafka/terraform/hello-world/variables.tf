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
    plan       = "business"
    flavor     = "db1-7"
    version    = "3.1"
  }
}

variable "access" {
  type = map(string)
  default = {
    ip = "xx.xx.xx.xx/32"
  }
}
