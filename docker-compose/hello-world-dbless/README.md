# Kong Gateway with Go Services Demo (DB-less Mode)

This project demonstrates a basic Kong API Gateway setup with two Go microservices using Docker Compose in DB-less mode.

## Project Structure

```
.
├── docker-compose.yml
├── kong.yml
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
mkdir -p kong-hello-world-dbless/service1 kong-hello-world-dbless/service2
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

## Testing the Services

Test the services using curl:
```bash
# Test service1
curl http://localhost:8000/service1

# Test service2
curl http://localhost:8000/service2
```

Expected responses:
- Service1: "Hello from Service 1!"
- Service2: "Hello from Service 2!"

## Port Configuration

- Kong Proxy: 8000 (API Gateway entry point)
- Kong Admin API: 8001 (Kong configuration endpoint)
- Service1: 8081 (direct access)
- Service2: 8082 (direct access)

## Components

- **Kong Gateway**: API Gateway in DB-less mode
- **Service1**: Go microservice returning "Hello from Service 1!"
- **Service2**: Go microservice returning "Hello from Service 2!"

## Key Differences from DB Version

- No PostgreSQL dependency
- Routes configured via declarative `kong.yml` file
- No need for migrations
- Faster startup time
- Configuration changes require container restart

## Notes

- In a production environment, only expose port 8000 (Kong Proxy) to the public
- Kong Admin API (8001) should be secured and not publicly accessible
- Direct service ports (8081, 8082) should only be accessible within the Docker network