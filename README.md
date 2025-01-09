MicroCraftGO: A Scalable Microservice in Golang

🚀 Overview:

This project leverages a microservice architecture, developed in Go (Golang), to deliver a high-performance, scalable, and maintainable application. It integrates gRPC for service-to-service communication and RabbitMQ for message brokering, ensuring efficient and reliable communication across services.

With Docker Compose managing the services, this application is designed for easy orchestration and local development. A Makefile has been added to simplify workflows and enhance task management.

🧩 Key Services:

**Logger Service**

   A centralized logging system, ensuring smooth tracking of application events.

   Database: MongoDB

**Authentication Service**

   Handles user authentication and security protocols.

   Database: PostgreSQL

**Mailer Service**:

   Manages email functionality for the application, ensuring communication via email.

   Testing: MailHog (local SMTP server for email testing)

**MailHog Service**

   A local SMTP server used to test and capture outgoing emails during development.

**Listener Service**

   A dedicated service for listening to RabbitMQ messages, enabling smooth communication between services.

**Broker Service**

   The primary entry point of the application, orchestrating requests and managing service coordination.

🌟 Features & Benefits:

 *Microservice Architecture*

  Designed for scalability and maintainability, allowing easy expansion.

*gRPC & RabbitMQ*

   Enables efficient inter-service communication with low latency and high throughput.

*Docker Compose*

   Simplifies orchestration and local setup for streamlined development processes.

*Independent Service Containers*

   Each service is containerized with its own dedicated Dockerfile for modular and isolated execution.

*Multiple Databases*

   MongoDB and PostgreSQL serve different needs, enabling flexible data management.

*MailHog Integration*

   Provides seamless email testing with a local SMTP server, making development efficient and reliable.

⚙️ Development & Deployment:

Project Status

✅ Development Phase Complete!

 Currently focusing on deployment, ensuring a smooth transition to production with robust scalability.

