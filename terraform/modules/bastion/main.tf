# eks-bastion-host/main.tf

resource "aws_instance" "bastion" {
  ami                    = var.ami_id
  instance_type          = "t3.micro"
  subnet_id              = var.public_subnet_id
  vpc_security_group_ids = [aws_security_group.bastion.id]
  key_name               = var.aws_key_name

  associate_public_ip_address = true

  tags = {
    Name = "eks-bastion"
  }
}

resource "aws_security_group" "bastion" {
  name        = "eks-bastion-sg"
  description = "Allow SSH to bastion from user IP"
  vpc_id      = var.aws_vpc_id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = var.allowed_ssh_cidr
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "bastion_public_ip" {
  value = aws_instance.bastion.public_ip
}

output "bastion_sg_id" {
    value = aws_security_group.bastion.id
}
