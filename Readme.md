# ITV Test Project

This is a test project for the company **ITV**. The project provides a REST API for user authentication and movie management.

## Technologies Used

- **Go** (Golang)
- **Gin** (HTTP framework)
- **GORM** (ORM for database management)
- **PostgreSQL** (Relational database)
- **UberFx** (Dependency injection)
- **Swagger** (API documentation)
- **Docker & Docker Compose** (Containerization)

# Project Structure

The project follows a clean architecture pattern with the following structure:
- **`config`** - Configuration files for environment variables and settings.
- **`database`** - Handles database connection and migrations.
- **`tools`** - Utility scripts and helper functions.
- **`customtypes`** - Defines custom data types used in the application.
- **`internal/entity`** - Contains domain models and entities.
- **`internal/usecase`** - Business logic layer, handling user and movie operations.
- **`internal/repository`** - Database access layer using GORM.
- **`external/ginapi`** - API layer (controllers) for handling HTTP requests.
- **`external/middleware`** - Authentication middleware for protected routes.

## Running the Project

To start the project, run the following script in the terminal:

```
./start_docker.sh
```

This script will set up the necessary services and run the application inside Docker containers.

## API Documentation

After starting the project, Swagger API documentation will be available at:

```
http://localhost:8080/swagger/index.html
```

This documentation provides an interactive interface for testing API endpoints.

## Endpoints

The API provides the following main endpoints:

### User Authentication
- `POST /user/register` – Register a new user
- `POST /user/auth` – Authenticate a user and receive a token
- `POST /user/logout` – Log out the user

### Movie Management
- `POST /movie` – Create a new movie (authorized users only)
- `GET /movie` – Get a list of all movies (authorized users only)
- `PUT /movie/:id` – Update movie details (authorized users only)
- `DELETE /movie/:id` – Delete a movie (authorized users only)

## Environment Variables

The application configuration is managed via environment variables defined in the `config` package. Ensure that the necessary variables are correctly set before running the project.

