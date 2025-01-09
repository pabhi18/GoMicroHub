MicroCarftGO: A Scalable Micro-service in golang 
🚀 Overview
This project leverages a microservice architecture, developed in Go (Golang), to deliver a high-performance, scalable, and maintainable application. It integrates gRPC for service-to-service communication and RabbitMQ for message brokering, ensuring efficient and reliable communication across services.

With Docker Compose managing the services, this application is designed for easy orchestration and local development. Makefile has been added to simplify workflows and enhance the ease of managing tasks.


Here’s a more polished and visually appealing version of your microservice project overview:

Microservice Project: A Scalable and Robust Solution
🚀 Overview
This project leverages a microservice architecture, developed in Go (Golang), to deliver a high-performance, scalable, and maintainable application. It integrates gRPC for service-to-service communication and RabbitMQ for message brokering, ensuring efficient and reliable communication across services.

With Docker Compose managing the services, this application is designed for easy orchestration and local development. Makefile has been added to simplify workflows and enhance the ease of managing tasks.

🧩 Key Services
1. Logger Service
A centralized logging system, ensuring smooth tracking of application events.

Database: MongoDB
2. Authentication Service
Handles user authentication and security protocols.

Database: PostgreSQL
3. Mailer Service
Manages email functionality for the application, ensuring communication via email.

Testing: MailHog (local SMTP server for email testing)
4. MailHog Service
A local SMTP server used to test and capture outgoing emails during development.

5. Listener Service
A dedicated service for listening to RabbitMQ messages, enabling smooth communication between services.

6. Broker Service
The primary entry point of the application, orchestrating requests and managing service coordination.


🌟 Features & Benefits
Microservice Architecture:
Designed for scalability and maintainability, allowing easy expansion.

gRPC & RabbitMQ:
Efficient inter-service communication with low latency and high throughput.

Docker Compose:
Simplified orchestration and local setup for smooth development processes.

Independent Service Containers:
Each service is containerized with its own dedicated Dockerfile for modular and isolated execution.

Multiple Databases:
MongoDB and PostgreSQL serve different needs, enabling flexibility in data management.

MailHog Integration:
Seamless email testing using MailHog, providing a local SMTP server for development.

⚙️ Development & Deployment
✅ Development Phase Complete!
Now, focus has shifted to deployment, ensuring a smooth transition to production with robust scalability.

🏗️ Running the Project Locally
Ensure Docker and Docker Compose are installed on your system.
Clone the repository to your local machine.
Navigate to the project directory.
Run the following command to start all services:

docker-compose up --build

🔧 Simplified Management with Makefile
Use the provided Makefile for a smoother workflow:

make build_up
