variable ovh {
    type    = map
    default = {
        endpoint           = "ovh-eu"
        application_key    = ""
        application_secret = ""
        consumer_key       = ""
    }
}

variable product {
    type    = map
    default = {
        project_id = ""
        region     = "DE"
        plan       = "essential"
        flavor     = "db1-7"
        version    = "4.0"
    }
}

variable access {
    type    = map
    default = {
        ip  = "xx.xx.xx.xx/32"
    }
}