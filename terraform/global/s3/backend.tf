terraform {
  backend "s3" {
    bucket  = "haikube"
    key     = "terraform-state/global/s3/terraform.tfstate"
    region  = "us-west-1"
    encrypt = true
  }
}