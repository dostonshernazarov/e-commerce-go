# E-Commerce Microservices Backend

This project is a backend implementation of an e-commerce platform built with Go using a microservices architecture. The services are designed to handle various operations like user management, product listings, order processing, comments, and notifications.

## Project Structure

The system is divided into the following services:

- **API Gateway**: Manages all incoming requests and routes them to the respective services.
- **User Service**: Handles user-related operations such as registration, login, and profile management.
- **Product Service**: Manages products, categories, and related information.
- **Order Service**: Responsible for managing orders, payments, and transactions.
- **Comment Service**: Allows users to comment on products.
- **Notification Service**: Sends notifications related to orders, products, and general platform updates.

## Technologies Used

- **Programming Language**: Go
- **Databases**: PostgreSQL, MongoDB
- **Message Brokers**: Kafka, Redis
- **Cloud Storage**: MinIO (S3 compatible)
- **Containerization**: Docker
- **Tracing and Monitoring**: Jaeger, OpenTelemetry

## Services Overview

### 1. API Gateway
- Handles routing of requests.
- Exposed at port `8085`.

### 2. User Service
- Manages user data and authentication.
- Exposed at port `8080`.
- Connected to PostgreSQL.

### 3. Product Service
- Manages products and categories.
- Exposed at port `8081`.
- Connected to MongoDB.

### 4. Order Service
- Handles order placement and payments.
- Exposed at port `8082`.
- Connected to PostgreSQL.

### 5. Comment Service
- Manages product reviews and user comments.
- Exposed at port `8084`.
- Connected to PostgreSQL.

### 6. Notification Service
- Sends notifications regarding orders, comments, and system updates.
- Exposed at port `8083`.
- Uses Redis for message queuing.

## Prerequisites

Before starting the project, make sure you have the following installed:

- Docker
- Docker Compose
- Go 1.19+
- PostgreSQL 14.x
- MongoDB

## Running the Project

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/dostonshernazarov/e-commerce-go.git
   cd e-commerce-go

2. **Build and Run with Docker Compose:** To run the entire system using Docker Compose, execute the following:
   ```bash
   docker-compose up --build
   
3. **Services:**

API Gateway: http://localhost:8085
Jaeger UI: http://localhost:16686 (for tracing)
MinIO: http://localhost:9000 (default user: `minioadmin`, password: `minioadmin`)

4. **Environment Variables:** Make sure to configure environment variables as per your system requirements. Some default values are provided in the `docker-compose.yml`.

5. **Migrations:** Database migrations for the PostgreSQL services (User, Order, Comment) are handled via the `migrate` service:
   ```bash
   docker-compose run migrate_user
   docker-compose run migrate_order
   docker-compose run migrate_comment

## Key Features
 - **Microservices:** Each service is independently deployable and scalable.
 - **RESTful API:** Every service follows REST principles, making it easy to interact with external clients.
 - **Message Queuing:** Kafka and Redis are used for asynchronous communication between services.
 - **Database:** PostgreSQL for relational data (Users, Orders, Comments), MongoDB for document-based data (Products).
 - **Tracing & Monitoring:** Integrated with OpenTelemetry and Jaeger for distributed tracing.

## Development
1. **Install dependencies:** Navigate to each service directory and run:
   ```bash
   go mod tidy

2. **Run services locally:** Each microservice can be run locally for development:
   ```bash
   go run main.go

## Future Improvements
 - Add gRPC for internal communication between microservices.
 - Implement more advanced notification mechanisms.
 - Extend tracing and logging.
 - Add unit tests and integration tests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.


