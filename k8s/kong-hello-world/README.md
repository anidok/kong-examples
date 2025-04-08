# Kong Hello World with Kubernetes

This project demonstrates a Kong API Gateway setup with two microservices in Kubernetes. The services are deployed in separate namespaces with Kong in the `kong-system` namespace and the applications in the `kong-hello-world` namespace.

## Project Structure

```
.
├── k8s/
│   ├── namespace.yaml
│   ├── kong-configmap.yaml
│   ├── kong-deployment.yaml
│   ├── service1-deployment.yaml
│   └── service2-deployment.yaml
├── service1/
│   ├── main.go
│   ├── Dockerfile
│   └── go.mod
└── service2/
    ├── main.go
    ├── Dockerfile
    └── go.mod
```

## Prerequisites

- Minikube installed
- kubectl installed
- Docker installed

## Deployment Steps

1. Start Minikube:
```bash
minikube start
```

2. Configure Docker to use Minikube's Docker daemon:
```bash
eval $(minikube docker-env)
```

3. Build and load the service images:
```bash
# Build and load service1
cd service1
docker build -t service1:latest .
minikube image load service1:latest

# Build and load service2
cd ../service2
docker build -t service2:latest .
minikube image load service2:latest
```

4. Apply Kubernetes manifests:
```bash
# Create namespaces
kubectl apply -f k8s/namespace.yaml

# Deploy Kong resources
kubectl apply -f k8s/kong-configmap.yaml
kubectl apply -f k8s/kong-deployment.yaml

# Deploy application services
kubectl apply -f k8s/service1-deployment.yaml
kubectl apply -f k8s/service2-deployment.yaml
```

5. Verify the deployments:
```bash
# Check Kong deployment
kubectl get pods -n kong-system
kubectl get svc -n kong-system

# Check application deployments
kubectl get pods -n kong-hello-world
kubectl get svc -n kong-hello-world
```

6. Access the Kong proxy:
```bash
minikube service kong-proxy -n kong-system
```

## API Endpoints

Once Kong is running, you can access the services through these endpoints:

- Service1 API1: `http://<kong-proxy-ip>/service1/api1`
- Service1 API2: `http://<kong-proxy-ip>/service1/api2`
- Service2 API1: `http://<kong-proxy-ip>/service2/api1`
- Service2 API2: `http://<kong-proxy-ip>/service2/api2`

## Architecture

- **Kong System Namespace**: Contains Kong API Gateway deployment and configuration
- **Kong Hello World Namespace**: Contains the application services
- **Service1**: Simple HTTP service running on port 8081
- **Service2**: Simple HTTP service running on port 8082

## Troubleshooting

If you encounter image pull errors:
1. Ensure you're using Minikube's Docker daemon: `eval $(minikube docker-env)`
2. Rebuild the images: `docker build -t <service>:latest .`
3. Load the images into Minikube: `minikube image load <service>:latest`
4. Delete the existing pod to force a new deployment: `kubectl delete pod -n kong-hello-world -l app=<service>`

## Cleanup

To remove all resources:
```bash
kubectl delete namespace kong-system
kubectl delete namespace kong-hello-world
```
