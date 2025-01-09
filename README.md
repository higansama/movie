# Movie App

This is a movie management application built with Go, using the Gin framework for the web server and GORM for database interactions. The application follows a clean architecture pattern.

## Features

- **Admin Functionality:**
  - Upload movies
  - Edit movies
  - Delete movies
  - Hide movies
  - List movies

- **User Functionality:**
  - View movie details
  - Rate movies

## Project Structure

```
movie-app
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── app
│   │   ├── app.go             # Application initialization and logic
│   │   └── routes.go          # Route definitions
│   ├── config
│   │   └── config.go          # Configuration settings
│   ├── controllers
│   │   ├── admin_controller.go # Admin controller for movie management
│   │   └── user_controller.go  # User controller for movie details and ratings
│   ├── models
│   │   └── movie.go            # Movie entity definition
│   ├── repositories
│   │   ├── movie_repository.go  # Movie repository interface
│   │   └── movie_repository_impl.go # Implementation of the movie repository
│   ├── services
│   │   ├── admin_service.go    # Business logic for admin operations
│   │   └── user_service.go     # Business logic for user operations
│   └── utils
│       └── utils.go            # Utility functions
├── go.mod                      # Module definition
├── go.sum                      # Module dependency checksums
└── README.md                   # Project documentation
```

## Controllers

Controllers in this project are responsible for handling HTTP requests and responses. They act as intermediaries between the client and the service layer, ensuring that the appropriate business logic is executed and the correct data is returned to the client.

- **Admin Controller (`admin_controller.go`):** Handles requests related to movie management, such as uploading, editing, deleting, and hiding movies.
- **User Controller (`user_controller.go`):** Handles requests related to viewing movie details and submitting ratings.

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd movie-app
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Configure the database connection in `internal/config/config.go`.

4. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

- Access the admin functionalities through the defined routes for managing movies.
- Users can view movie details and submit ratings via the user routes.

## License

This project is licensed under the MIT License.