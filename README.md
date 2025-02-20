# MicroCraftGo: A Scalable Microservice Architecture in Golang

## Overview

MicroCraftGo is a microservice-based application developed in Go (Golang) designed for high performance, scalability, and maintainability. It utilizes gRPC for efficient service-to-service communication and RabbitMQ for message brokering. The application is containerized using Docker and orchestrated with Docker Compose for seamless local development and deployment.

## Features

- **Logger Service**: Centralized logging using MongoDB for tracking application events.
- **Authentication Service**: Manages user authentication with PostgreSQL as the database.
- **Mailer Service**: Handles email functionalities using MailHog for testing.
- **Broker Service**: Acts as the primary entry point, coordinating requests and service interactions.
- **Listener Service**: Listens to RabbitMQ messages for inter-service communication.
- **gRPC & RabbitMQ Integration**: Ensures low-latency, high-throughput communication.

## Prerequisites

- **Docker**: Ensure Docker is installed and running on your machine.
- **Docker Compose**: Required for orchestrating the services.
- **Go**: Install Go for local development and testing.

## Getting Started Locally

 **Clone the Repository**:
   ```bash
   git clone https://github.com/sanjaygupta972004/MicroCraftGo.git
   cd MicroCraftGo
   ```

 **Use Docker Compose to build and start all services**:
   
 ```bash
   docker-compose up --build
```

