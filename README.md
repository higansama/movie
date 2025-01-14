# Movie App

This is a movie management application built with Go, using the Gin framework for the web server and GORM for database interactions. The application follows a clean architecture pattern.

## Features

- **Admin Functionality:**
  - Upload movies
  - Edit movies
  - Delete movies
  - Hide movies
  - List movies
  - Login

- **User Functionality:**
  - View movie details
  - Rate movies
  - Register
  - Login
  - Searching Movie

## Project Structure


```
movie-app
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── core                   # Abstraction of functionality in the project. It should move to outer folder
│   │   ├── repositories             # Main Repositories Interface
|   │   │   ├── actors               # For Actor models functionality
|   │   │   ├── castings             # For Casting models functionality
|   │   │   ├── genre                # For Genre models functionality
|   │   │   ├── movie                # For Movie models functionality
|   │   │   ├── watch                # For Watch models functionality
│   │   ├── reqres                   # Request And Response, in other words is JSON DTO, just in case you want to using gRPC, create protoc in other folder and assign to this model
│   │   ├── services                 # Application initialization and logic
|   │   │   ├── admin                # Admin abstraction functionality
|   │   │   ├── user                 # User abstraction functionality
│   ├── entrypoint
│   │   ├── admin.go                 # Admin module initialization and logic
│   │   └── user.go                  # User module initialization and logic
│   ├── config
│   │   └── config.go          # Configuration settings
│   ├── controllers
│   │   ├── admin_controller.go # Admin controller for movie management
│   │   └── user_controller.go  # User controller for movie details and ratings
│   ├── models
│   │   ├── movie.go            # Movie entity definition, and some models about the movie
│   │   ├── users.go            # Users entity definition
│   ├── repositories
│   │   ├── cast_implementation.go           # Casting repository implementation
│   │   ├── genre_implementation.go          # Genre repository implementation
│   │   ├── movie_implementation.go          # Movie repository implementation
│   │   ├── user_implementation.go           # User repository implementation
│   │   └── wathc_implementation.go          # WatchHistory repository implementation
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

4. Migration
   Migrate the models:
   ```
   go run cmd/main.go migrate
   ```
   -------------------------------------------------------------------------------------------------------------------------
   Create admin: 
   ```
   go run cmd/main.go createadmin
   ```
   -------------------------------------------------------------------------------------------------------------------------
   Make a seed
   ```
   go run cmd/main.go seed-actor
   go run cmd/main.go seed-genre
   ```

5. Run the application:
   ```
   go run cmd/main.go runserver
   ```


## Usage

- Access the admin functionalities through the defined routes for managing movies.
- Users can view movie details and submit ratings via the user routes.

## License

This project is licensed under the MIT License.