# GO + GORM + PostgreSQL

This is the backend service for HelperHub, built using Go, Gin, and GORM. It provides APIs for user management and other functionalities.

## Prerequisites

- Go 1.16 or later
- PostgreSQL

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/HelperHub.git
    cd HelperHub/backend
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Set up PostgreSQL:

    - Ensure PostgreSQL is installed and running.
    - Create a PostgreSQL user and database.

4. Update the database connection string in `routes.go`:

    ```go
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
    ```

## Running the Application

1. Initialize the database and start the server:

    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:8080`.

## API Endpoints

### Ping

- **URL:** `/ping`
- **Method:** `GET`
- **Description:** Returns a pong message to check if the server is running.

    ```sh
    curl -X GET http://localhost:8080/ping
    ```

### Create User

- **URL:** [users](http://_vscodecontentref_/1)
- **Method:** `POST`
- **Description:** Creates a new user.

    ```sh
    curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{
      "Username": "testuser",
      "Email": "testuser@example.com",
      "Password": "password123",
      "Volunteer": true,
      "Voluntee": false,
      "Category": "General"
    }'
    ```

## Project Structure

- `main.go`: Entry point of the application.
- `routes.go`: Contains route definitions and handlers.
- `models.go`: Contains database models.

## License

This project is licensed under the MIT License.