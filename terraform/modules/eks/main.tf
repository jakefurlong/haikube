# eks-spot-cluster/main.tf

provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.1.0"

  name = var.aws_vpc_name
  cidr = var.aws_vpc_cidr

  azs             = slice(data.aws_availability_zones.available.names, 0, 2)
  public_subnets  = var.aws_public_subnets
  private_subnets = var.aws_private_subnets

  enable_nat_gateway = true
  single_nat_gateway = true
}

data "aws_availability_zones" "available" {}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  version         = "20.36.0"
  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version

  subnet_ids         = module.vpc.private_subnets
  vpc_id             = module.vpc.vpc_id
  enable_irsa        = true
  cluster_endpoint_public_access         = true

  cluster_addons = {
    coredns = {
      most_recent                 = true
      resolve_conflicts_on_create = "OVERWRITE"
      configuration_values        = jsonencode(local.coredns_config)
    }
    kube-proxy = {
      most_recent = true
    }
    vpc-cni = {
      most_recent              = true
      service_account_role_arn = aws_iam_role.vpc_cni.arn
    }
  }

  enable_cluster_creator_admin_permissions = true

  eks_managed_node_groups = {
    single_node = {
      desired_size = 1
      min_size     = 1
      max_size     = 1

      instance_types = var.aws_instance_types
      capacity_type  = "SPOT"

      labels = {
        lifecycle = "Ec2Spot"
      }

      tags = {
        Name = "eks-single-spot-node"
      }
    }
  }
}



#module "aws_auth" {
#  source  = "terraform-aws-modules/eks/aws//modules/aws-auth"
#  version = "20.8.4"
#
#  aws_auth_users = [
#    {
#      userarn  = "arn:aws:iam::443374376889:user/jfurlong"
#      username = "jfurlong"
#      groups   = ["system:masters"]
#    }
#  ]
#
#  depends_on = [module.eks]
#}

output "vpc_id" {
  description = "The ID of the created VPC"
  value       = module.vpc.vpc_id
}

output "public_subnet_id" {
  description = "The first public subnet ID from the VPC"
  value       = module.vpc.public_subnets[0]
}

output "cluster_security_group_id" {
  value = module.eks.cluster_security_group_id
}