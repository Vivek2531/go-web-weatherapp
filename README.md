# Weather Application

A simple, lightweight web application built with Go to demonstrate cloud-native development and deployment practices.

## Overview

This is a minimal web application that provides a weather information interface. It's designed to showcase containerization, Kubernetes deployment, and modern development workflows.

## Technology Stack

- **Backend**: Go (Golang) - net/http package
- **Frontend**: HTML5, CSS3, JavaScript
- **Container**: Docker (multi-stage builds)
- **Deployment**: Kubernetes (AWS EKS)

## Quick Start

### Local Development

```bash
go run main.go
```

Application runs on `http://localhost:8000`

### Docker

```bash
# Build
docker build -t weather-app .

# Run
docker run -p 8000:8000 weather-app
```

### Kubernetes

```bash
kubectl apply -f k8s/manifests/
```

## Project Structure

```
.
├── main.go              # Application server
├── main_test.go         # Tests
├── go.mod               # Go dependencies
├── Dockerfile           # Container build
├── static/              # HTML files
│   ├── home.html
│   └── about.html
└── k8s/                 # Kubernetes manifests
    └── manifests/
        ├── deployment.yaml
        ├── service.yaml
        └── ingress.yaml
```

## Features

- Lightweight and fast
- Containerized deployment
- Kubernetes-ready
- Simple and clean interface

## Configuration

- Port: 8000
- Base image: distroless/base
- Build image: golang:1.25.4-alpine3.22

## License

MIT


