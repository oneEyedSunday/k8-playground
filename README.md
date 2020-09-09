# k8-playground
Playground for K8

## Start minikube (local)
minikube start --driver=hyperkit

### Docker context
eval $(minikube docker-env)

### Stop Minikube
minikube stop


## Make scripts executable
chmod a+x simple/*.sh
