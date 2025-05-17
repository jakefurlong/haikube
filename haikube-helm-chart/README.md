You need a secret for AWS creds

`kubectl create secret generic aws-credentials --from-literal=AWS_ACCESS_KEY_ID=<STRING>> --from-literal=AWS_SECRET_ACCESS_KEY=<STRING> --from-literal=AWS_REGION=us-west-1`

```
minikube start --driver=docker --addons=ingress

mac and linux...  
eval $(minikube docker-env)

windows
minikube image load haikube-backend:latest
minikube image load haikube-frontend:latest

helm install haikube ./haikube-helm-chart # or...
helm upgrade haikube ./haikube-helm-chart

```

