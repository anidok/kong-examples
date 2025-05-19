# Kong Hello World with Kubernetes and JWT Authentication

This project demonstrates a Kong API Gateway setup with two microservices in Kubernetes, including JWT authentication. The services are deployed in separate namespaces with Kong in the `kong-system` namespace and the applications in the `kong-hello-world` namespace.

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
- Node.js and npm installed

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

## JWT Authentication Setup

1. Install the required Node.js dependencies:
   ```bash
   npm init -y
   npm install jsonwebtoken
   ```

2. Create a JWT helper script (`generate-jwt.js`):
   ```js
   const jwt = require('jsonwebtoken');

   const secret = 'some-key';  // Same secret as in Kong config

   function generateToken() {
       const payload = {
           iss: 'kong-jwt-auth',
           exp: Math.floor(Date.now() / 1000) + 3600,
           sub: 'user123'
       };
       const token = jwt.sign(payload, secret);
       console.log('Generated JWT Token:');
       console.log(token);
   }

   function verifyToken(token) {
       try {
           const decoded = jwt.verify(token, secret);
           console.log('Token is valid. Decoded payload:');
           console.log(JSON.stringify(decoded, null, 2));
       } catch (err) {
           console.error('Token verification failed:', err.message);
       }
   }

   // Handle command line arguments
   const command = process.argv[2];
   const token = process.argv[3];

   if (command === 'generate') {
       generateToken();
   } else if (command === 'verify' && token) {
       verifyToken(token);
   } else {
       console.log('Usage:');
       console.log('  node generate-jwt.js generate');
       console.log('  node generate-jwt.js verify <token>');
       process.exit(1);
   }
   ```

3. Update the Kong configuration (`k8s/kong-configmap.yaml`) to include a consumer and JWT secret:
   ```yaml
   consumers:
     - username: jwt-user
       custom_id: user123

   acls:
     - consumer: jwt-user
       group: user123-group

   jwt_secrets:
     - consumer: jwt-user
       key: kong-jwt-auth
       secret: some-key
   ```

4. Re-apply the updated Kong configuration:
   ```bash
   kubectl apply -f k8s/kong-configmap.yaml
   ```

5. Restart the Kong pod to load the new configuration:
   ```bash
   kubectl delete pod -n kong-system -l app=kong
   ```

6. Wait for the Kong pod to restart:
   ```bash
   kubectl get pods -n kong-system
   ```

7. Port-forward the Kong service to make it accessible locally:
   ```bash
   kubectl port-forward -n kong-system svc/kong-proxy 8000:80
   ```

## Using JWT Authentication

1. Generate a JWT token:
   ```bash
   node generate-jwt.js generate
   ```

2. Use the token in your API requests:
   ```bash
   curl -v -H "Authorization: Bearer <your-token>" http://localhost:8000/service1/api1
   ```

3. Verify a token:
   ```bash
   node generate-jwt.js verify <your-token>
   ```

## Configuration Details

The JWT plugin is configured with the following settings:
- `key_claim_name`: "iss" (issuer)
- `claims_to_verify`: ["exp"] (expiration time)

The token payload includes:
- `iss`: Token issuer
- `exp`: Expiration time (1 hour from generation)
- `sub`: Subject (user identifier)

## Security Notes

1. Always use HTTPS in production
2. Keep your secret key secure and rotate it periodically
3. Consider implementing token refresh mechanisms
4. Set appropriate token expiration times

## Troubleshooting

If you encounter authentication issues:
1. Verify the token hasn't expired
2. Check the token format in the Authorization header
3. Ensure the Kong JWT plugin is properly configured
4. Check Kong logs for detailed error messages

## Cleanup

To remove all resources:
```bash
kubectl delete namespace kong-system
kubectl delete namespace kong-hello-world
```

## Kong JWT and ACL Behavior

- **JWT Plugin:**  
  Kong uses the `iss` claim to map the JWT to a consumer. The `sub` claim is not used for access control.

- **ACL Plugin:**  
  Kong checks if the mapped consumer is in the allowed group. The `sub` claim is not validated by Kong.

- **Subject-Based Access:**  
  If you need to enforce access based on the `sub` claim, you have two options:
  1. **Enforce in your upstream service:**  
     Your backend should check the `sub` claim in the JWT payload.
  2. **Use one consumer per subject:**  
     Create a Kong consumer for each subject, issue JWTs with a unique `iss` per user, and tie each JWT secret to a specific consumer.

Example for multiple users:
```yaml
consumers:
  - username: user123
    custom_id: user123
  - username: user124
    custom_id: user124

acls:
  - consumer: user123
    group: user123-group
  - consumer: user124
    group: user124-group

jwt_secrets:
  - consumer: user123
    key: user123-key
    secret: user123-secret
  - consumer: user124
    key: user124-key
    secret: user124-secret
```

Then, issue JWTs with `iss: user123-key` for user123, and so on.
