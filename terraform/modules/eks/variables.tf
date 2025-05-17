variable "aws_region" {
  description = "Default AWS region"
  type = string
}

variable "aws_vpc_name" {
  description = "name of the EKS VPC"
  type = string
}

variable "aws_vpc_cidr" {
  description = "CIDR range of the VPC"
  type = string
}

variable "aws_public_subnets" {
  description = "Public subnets of the EKS VPC"
  type = list(string)
}

variable "aws_private_subnets" {
  description = "Private subnets of the EKS VPC"
  type = list(string)
}

variable "cluster_name" {
  description = "Name of the Kubernetes cluster"
  type = string
}

variable "cluster_version" {
  description = "Version of the EKS cluster"
  type = string
}

variable "aws_instance_types" {
  description = "Instance types for EKS nodes"
  type = list(string)
}