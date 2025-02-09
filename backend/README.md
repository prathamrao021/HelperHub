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

4. Update the database connection string in `main.go`:

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

### Create User

- **URL:** `/users/create`
- **Method:** `POST`
- **Description:** Creates a new user.

    ```sh
    curl -X POST http://localhost:8080/users/create \
    -H "Content-Type: application/json" \
    -d '{
      "Username": "testuser",
      "Email": "testuser@example.com",
      "Password": "password123",
      "Volunteer": true,
      "Voluntee": false,
      "Category": ["General", "Technology", "Health"]
    }'
    ```

### Update User

- **URL:** `/users/update/{username}`
- **Method:** `PUT`
- **Description:** Updates an existing user.

    ```sh
    curl -X PUT http://localhost:8080/users/update/testuser \
    -H "Content-Type: application/json" \
    -d '{
      "Username": "testuser",
      "Email": "newemail@example.com",
      "Password": "newpassword123",
      "Volunteer": true,
      "Voluntee": false,
      "Category": ["General", "Technology", "Health"]
    }'
    ```

### Delete User

- **URL:** `/users/delete/{username}`
- **Method:** `DELETE`
- **Description:** Deletes an existing user.

    ```sh
    curl -X DELETE http://localhost:8080/users/delete/testuser
    ```


## Swagger Documentation

### Generating Swagger Documentation

1. Install the necessary packages:

    ```sh
    go get -u github.com/swaggo/swag/cmd/swag
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/files
    ```

2. Generate the Swagger documentation:

    ```sh
    swag init
    ```

3. Import the generated docs in your `main.go` file:

    ```go
    import _ "github.com/yourusername/HelperHub/docs"
    ```

4. Add the Swagger setup to your `main.go` file:

    ```go
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    ```

5. Add the Swagger route in your `main.go` file:

    ```go
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    ```

### Accessing Swagger Documentation

Once the server is running, you can access the Swagger UI at:

    ```
    http://localhost:8080/swagger/index.html
    ```

## Project Structure

- `main.go`: Entry point of the application.
- `routes.go`: Contains route definitions and handlers.
- `models.go`: Contains database models.

## License

This project is licensed under the MIT License.