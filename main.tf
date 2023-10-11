terraform {
  required_providers {
    terratowns = {
      source  = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
}



# https://registry.terraform.io/providers/hashicorp/random/latest/docs


#https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
#terraform {
#cloud {
#  organization = "Heather"
#  workspaces {
#    name = "terra-house-8"
#  }
#}

#}

provider "terratowns" {
  endpoint  = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid
  token = var.terratowns_access_token

}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.teacherseat_user_uuid
  bucket_name = var.bucket_name
  index_html_filepath = var.index_html_filepath
  error_html_filepath = var.error_html_filepath
  content_version = var.content_version
  assets_path = var.assets_path
}

resource "terratowns_home" "Home" {
  name        = "How to learn Astrology in 2023!"
  description = <<DESCRIPTION
  This is how I would learn Astrology
  DESCRIPTION
  domain_name = module.terrahouse_aws.cloudfront_url
  #domain_name = "3ghdhr.cloudfront.net"
  town            = "missingo"
  content_version = 1
}
