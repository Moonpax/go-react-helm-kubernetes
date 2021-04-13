# Go React Helm Kubernetes

## Stack:

- Docker
- Kubernetes
- Helm 3 Charts
- Golang with gin-gonic
- React app

## Project

- Server folder: Golang server with Elon Musk SpaceX Api return last launches
- Client folder: React small app shows last launches with Failure description
- k8s folder: Chart-Api settings with image name and replicaCount: 3
- Kubernetes service LoadBalancer with 3 pods

## Commands for build and push docker to Docker https://hub.docker.com/

### server:

`$ docker build -t moonpax/go-api .`

`$ docker push moonpax/go-api`

### client:

`$ docker build -t moonpax/react-client .`
`$ docker push -t moonpax/react-client`

### helm3 deploy command:

`$ helm install go-api k8s/Chart-Api/`
`$ helm uninstall go-api`

### Additional k8s commans:

`kubectl describe service`
`kubectl get replicaset`
`kubectl get pods`
`kubectl get svc`
`helm ls`