Openleaf Backend Assignment – Order Shipping API (Golang)

This project is a Golang backend service that simulates an order shipping system.
It demonstrates clean architecture, concurrency, failure handling, and observability.

The system supports:

Creating an order with a selected carrier

Fetching shipping rates concurrently from multiple carriers

Polling shipment tracking updates asynchronously

Tech Stack

Go (Golang)

net/http

slog (logging)

lumberjack (log rotation)

In-memory storage

Architecture Overview

The project follows a layered architecture:

cmd/ → Application entry point
routes/ → HTTP routing
controllers/ → Request handling & validation
services/ → Business logic
repository/ → In-memory persistence
models/ → Domain models & API contracts
logger/ → Custom text logger (slog + lumberjack)
lib/helpers/ → Carrier simulations

Each layer has a single responsibility and minimal coupling.

API Endpoints

Create Order
POST /api/v1/order/create

Validates input, checks order uniqueness, simulates carrier serviceability,
generates a tracking ID, and stores the order.

Rate Calculator (Concurrent)
POST /api/v1/order/rates

Fetches shipping rates from multiple simulated carriers concurrently using
goroutines and channels.

Only serviceable carriers are returned.
Carrier failures or delays do not affect the overall response.

Tracking Poller
POST /api/v1/tracking/poll

Polls simulated carrier tracking APIs and inserts only new tracking events.
Duplicate events are avoided, and carrier failures do not crash the system.

Logging

Custom text logger built using slog and lumberjack

Human-readable log format

Includes timestamp, log level, message, source file, and function (for errors)

Logs are written to:

logs/app.log

___________________________________________________________________________________________


How to Run

Prerequisites:

Go 1.21 or higher

Steps:

git clone https://github.com/mgsquare/openleaf_assignment.git

cd openleaf_assignment
go mod tidy
go run cmd/main.go

The server starts at:

http://localhost:8080

Sample Requests

Create Order
POST /api/v1/order/create

{
"order_id": "ORD123",
"carrier": "bluedart",
"pincode": "560001"
}

Calculate Rates
POST /api/v1/order/rates

{
"pincode": "560001",
"weight": 500
}

Poll Tracking
POST /api/v1/tracking/poll

{
"tracking_id": "TRK-12345"
}

Design Choices

In-memory storage keeps the system lightweight and easy to run

Carrier APIs are simulated to focus on system behavior rather than integrations

Concurrency is used where it provides real value (rate calculation)

Logging is optimized for human readability instead of machine parsing

Future Improvements

Retries with exponential backoff for carrier APIs

Per-carrier timeouts and circuit breakers

Persistent database (Postgres or MySQL)

Background workers for tracking polling

Message queues for async processing

Metrics and distributed tracing

Authentication and rate limiting

Comprehensive unit and concurrency tests

CI/CD pipeline integration

Notes

The implementation prioritizes correctness, clarity, and extensibility.
The current design allows production-grade features to be added without major refactoring.