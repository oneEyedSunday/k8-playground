## Starting up

Start up local kubernetes, using minikube and the docker driver

```sh
minikube start --driver=docker
```

## Resources

This is a Kubernetes manifest deploying an application infrastructure containing the following:

- API Services
- 2 PostgresSQL databases
- Ingress controller


We are not concerned about persisting data for the databases (this is not a production installation).


### Deploy instructions

Create namespaces 
```sh
kubectl create namespace basic-app
```

Set the kubectl context to our app namespace
```sh
kubectl config set-context --current --namespace=basic-app
```

Deploy database(s)

```sh
kubectl apply -f dbs/ 
```

Inspect your database

```sh
# log in to your database shell
psql -U app_user -d database1

# inspect database schemas, tables and see other databases
\d
\dt
\l
```

TODO use helm to automate some of these db configuration

Deploy api services

Deploy service to registry (just use https://ttl.sh/ for now)


IMAGE_NAME="oneeyedsunday_api1_$(uuidgen)"

```sh
IMAGE_NAME="oneeyedsunday_api1"                          
docker build -t ttl.sh/${IMAGE_NAME}:1h .                                
docker push ttl.sh/${IMAGE_NAME}:1h
```

```sh
kubectl apply -f services/
```