

module "eks" {
  source = "../../modules/eks"

  aws_region          = "us-west-1"
  aws_vpc_name        = "eks-vpc"
  aws_vpc_cidr        = "10.0.0.0/16"
  aws_public_subnets  = ["10.0.1.0/24", "10.0.2.0/24"]
  aws_private_subnets = ["10.0.3.0/24", "10.0.4.0/24"]
  cluster_name        = "haikube-production"
  cluster_version     = "1.32"
  aws_instance_types  = ["t3.medium"]
}

module "bastion" {
  source = "../../modules/bastion"

  ami_id           = "ami-07706bb32254a7fe5" # us-west-1 AL2023
  public_subnet_id = module.eks.public_subnet_id
  allowed_ssh_cidr = ["47.33.30.15/32"]
  aws_key_name     = "haikube-key"
  aws_vpc_id       = module.eks.vpc_id

  depends_on = [module.eks]
}

resource "aws_security_group_rule" "allow_bastion_to_eks" {
  type                     = "ingress"
  from_port                = 443
  to_port                  = 443
  protocol                 = "tcp"
  source_security_group_id = module.bastion.bastion_sg_id
  security_group_id        = module.eks.cluster_security_group_id
  description              = "Allow bastion to access EKS API"
}