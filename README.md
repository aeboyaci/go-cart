## Description
Containerized backend implementation for a cart application along with Integration & Unit tests

## Information

**Features**
- Containerized with Docker & Docker-Compose
- Authentication with JWT
- Role Based Authorization
- Dependency Injection
- Mocking
- Integration Tests (`**/repository_test.go`)
- Unit Tests (`**/service_test.go`)
- Middleware Tests (`pkg/middlewares/authentication_test.go`)
- Swagger API Documentation

**Programming Language:** Go (1.19)

**Framework:** Echo

**Database:** PostgreSQL

**Code Separation**

Responsibility separation is applied in files as follows:

- *router.go:* registering routers
- *controller.go*: calling the related service functions and returning responses
- *service.go:* business logic
- *repository.go:* database queries
- *params.go:* holding Objects that are being used in the requests
- *dto.go:* holding Data Transfer Objects

**Author:** Ahmet Eren BOYACI
