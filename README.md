# Go Web Weather App

A simple weather application built with Go, deployed on Kubernetes with automated CI/CD.

![Home Page](docs/images/home-page.png)

![About Page](docs/images/about-page.png)

## What is this?

A lightweight web application that displays weather information. It demonstrates a complete DevOps workflow from development to production deployment with HTTPS.

**Live Demo:** https://cloudopsprojects.com

## Technologies Used

- **Go 1.22** - Backend application
- **Docker** - Containerization with multi-stage builds
- **Kubernetes** - Container orchestration on AWS EKS
- **Helm** - Kubernetes package manager
- **NGINX Ingress** - Load balancing and routing
- **cert-manager** - Automatic SSL certificates from Let's Encrypt
- **GitHub Actions** - CI/CD automation

## Local Development

```bash
# Run locally
go run main.go

# Visit http://localhost:8000/home
```

## Build & Deploy

```bash
# Build Docker image
docker build -t go-web-app .

# Deploy to Kubernetes
kubectl apply -f k8s/manifests/

# Or use Helm
helm install go-web-app ./helm/go-web-app
```

## CI/CD

The GitHub Actions pipeline automatically:
- Builds and tests the application
- Creates Docker images and pushes to DockerHub
- Updates Helm charts with new versions

Every push to `main` triggers the full pipeline.

## Project Structure

```
├── main.go                 # Application code
├── static/                 # HTML files
├── Dockerfile              # Container configuration
├── k8s/manifests/          # Kubernetes deployment files
├── helm/go-web-app/        # Helm chart
└── .github/workflows/      # CI/CD pipeline
```

## License

MIT

## Author

Vivek Inturi
