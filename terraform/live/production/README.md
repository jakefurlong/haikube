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