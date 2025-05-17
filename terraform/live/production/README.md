```
terraform apply

aws eks update-kubeconfig --name haikube-production --region us-west-1
aws eks --region us-west-1 update-kubeconfig --name haikube-production

# install kubectl
curl -O https://s3.us-west-2.amazonaws.com/amazon-eks/1.32.3/2025-04-17/bin/linux/amd64/kubectl
chmod +x ./kubectl
mkdir -p $HOME/bin && cp ./kubectl $HOME/bin/kubectl && export PATH=$HOME/bin:$PATH
echo 'export PATH=$HOME/bin:$PATH' >> ~/.bashrc

```

## Cloudformation / AWS CLI version

aws cloudformation create-stack --region us-west-1 --stack-name my-eks-vpc-stack --template-url https://s3.us-west-2.amazonaws.com/amazon-eks/cloudformation/2020-10-29/amazon-eks-vpc-private-subnets.yaml

aws iam create-role --role-name myAmazonEKSClusterRole --assume-role-policy-document file://"eks-cluster-role-trust-policy.json"

aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/AmazonEKSClusterPolicy --role-name myAmazonEKSClusterRole

Create cluster manually following doc: https://docs.aws.amazon.com/eks/latest/userguide/getting-started-console.html

aws eks update-kubeconfig --region us-west-1 --name my-cluster
            
aws iam create-role --role-name myAmazonEKSNodeRole --assume-role-policy-document file://"node-role-trust-policy.json"

aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy --role-name myAmazonEKSNodeRole
aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly --role-name myAmazonEKSNodeRole
aws iam attach-role-policy --policy-arn arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy --role-name myAmazonEKSNodeRole

Create the node group manually following doc: https://docs.aws.amazon.com/eks/latest/userguide/getting-started-console.html

Add policy to Node role manually to access ECR:

```
{
  "Effect": "Allow",
  "Action": [
    "ecr:GetAuthorizationToken",
    "ecr:BatchCheckLayerAvailability",
    "ecr:GetDownloadUrlForLayer",
    "ecr:BatchGetImage"
  ],
  "Resource": "*"
}
```

aws ecr get-login-password --region us-west-1 | docker login --username AWS --password-stdin 443374376889.dkr.ecr.us-west-1.amazonaws.com

docker tag haikube-backend 443374376889.dkr.ecr.us-west-1.amazonaws.com/haikube-backend:latest

Create a k8s secret