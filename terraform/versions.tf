terraform {
  required_providers {
    example = {
      version = ">= 1.0.0, <= 2.0.0"
      source  = "terraform-example.com/exampleprovider/example"
    }
  }
  required_version = ">= 1.1.0"
}

provider "example" {
  # ...
}
