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



## Jeager
   ```bash
   package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	service     = "example-api"
	environment = "development"
	id          = 1
)

type Message struct {
	Text string `json:"text"`
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

// swagger
func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tracer := otel.Tracer(service)

	_, span := tracer.Start(ctx, "helloHandler")
	defer span.End()

	time.Sleep(50 * time.Millisecond)

	span.SetAttributes(attribute.String("handler", "hello"))

	msg := Message{Text: "Hello, World!"}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	handler := otelhttp.NewHandler(http.HandlerFunc(helloHandler), "hello")

	http.Handle("/hello", handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


```

## Run Jeager with Docker
```bash
 docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.25
