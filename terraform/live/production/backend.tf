terraform {
  backend "s3" {
    bucket  = "haikube"
    key     = "terraform-state/live/production/terraform.tfstate"
    region  = "us-west-1"
    encrypt = true
  }
}