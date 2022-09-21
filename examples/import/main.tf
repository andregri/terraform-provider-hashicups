terraform {
  required_providers {
    hashicups = {
      version = "0.2"
      source  = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
  username = "education"
  password = "test123"
  url      = "http://192.168.1.203:19090"
}

resource "hashicups_order" "sample" {}

output "sample_order" {
  value = hashicups_order.sample
}

