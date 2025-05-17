variable "ami_id" {
  description = "AMI ID"
  type = string
}

variable "public_subnet_id" {
  description = "Public Subnet ID"
  type = string
}

variable "allowed_ssh_cidr" {
  description = "Allows ingress CIDR range"
  type = list(string)
}

variable "aws_key_name" {
  description = "AWS Key Pair Name"
  type = string
}

variable "aws_vpc_id" {
  description = "VPC ID"
  type = string
}