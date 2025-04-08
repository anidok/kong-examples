# Kong Gateway with Go Services Demo

This project demonstrates a basic Kong API Gateway setup with two Go microservices using Docker Compose.

## Project Structure

```
.
├── docker-compose.yml
├── README.md
├── service1
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
└── service2
    ├── Dockerfile
    ├── go.mod
    └── main.go
```

## Setup Instructions

1. Create the directory structure:
```bash
mkdir service1 service2
```

2. Initialize Go modules:
```bash
cd service1
go mod init service1
cd ../service2
go mod init service2
```

3. Create all the service files as provided in this repository

4. Start the services:
```bash
docker-compose up -d
```

5. Configure Kong routes using the Admin API:
```bash
# Create service1 in Kong
curl -i -X POST http://localhost:8001/services \
  --data name=service1 \
  --data url='http://service1:8081'

# Create routes for service1
curl -i -X POST http://localhost:8001/services/service1/routes \
  --data paths[]=/service1/api1 \
  --data strip_path=true \
  --data name=service1-api1

curl -i -X POST http://localhost:8001/services/service1/routes \
  --data paths[]=/service1/api2 \
  --data strip_path=true \
  --data name=service1-api2

# Create service2 in Kong
curl -i -X POST http://localhost:8001/services \
  --data name=service2 \
  --data url='http://service2:8082'

# Create routes for service2
curl -i -X POST http://localhost:8001/services/service2/routes \
  --data paths[]=/service2/api1 \
  --data strip_path=true \
  --data name=service2-api1

curl -i -X POST http://localhost:8001/services/service2/routes \
  --data paths[]=/service2/api2 \
  --data strip_path=true \
  --data name=service2-api2
```

## Testing the Services

Test the services using curl:
```bash
# Test service1 endpoints
curl http://localhost:8000/service1/api1
curl http://localhost:8000/service1/api2

# Test service2 endpoints
curl http://localhost:8000/service2/api1
curl http://localhost:8000/service2/api2
```

Expected responses:
- Service1 API1: "Hello from Service 1 - API 1!"
- Service1 API2: "Hello from Service 1 - API 2!"
- Service2 API1: "Hello from Service 2 - API 1!"
- Service2 API2: "Hello from Service 2 - API 2!"

## Port Configuration

- Kong Proxy: 8000 (API Gateway entry point)
- Kong Admin API: 8001 (Kong configuration endpoint)
- Service1: 8081 (direct access)
- Service2: 8082 (direct access)

## Components

- **Kong Gateway**: API Gateway with PostgreSQL backend
- **PostgreSQL**: Database for Kong configuration
- **Service1**: Go microservice with two API endpoints
- **Service2**: Go microservice with two API endpoints

## Notes

- In a production environment, only expose port 8000 (Kong Proxy) to the public
- Kong Admin API (8001) should be secured and not publicly accessible
- Direct service ports (8081, 8082) should only be accessible within the Docker network
