# kubieous

Internal cluster monitoring

## Running locally with minikube

- $ cd kubieous
- $ minikube start
- $ minikube dashboard
- $ eval $(minikube docker-env)
- $ docker build -t kubieous:v1 .
- $ kubectl run kubieous --image=kubieous:v1 --port=8080
- You'll also want to run an image that you can test on

## Testing HPA
- You need to enable heapster `$ minikube addons enable heapster`
