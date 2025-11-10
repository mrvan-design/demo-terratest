terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.66"
    }
  }
}

provider "aws" {
  region                      = var.aws_region
  access_key                  = "test"
  secret_key                  = "test"
  s3_use_path_style           = true
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  endpoints {
 
    ec2 = "http://127.0.0.1:4566"
    s3  = "http://127.0.0.1:4566"
    sts = "http://127.0.0.1:4566" # thêm dòng này cho S3
  }
}
