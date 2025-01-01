# Go User CRUD API

This is a simple user CRUD (Create, Read, Update, Delete) application built with Go (Golang). It provides a basic set of endpoints to manage users in a database. This project demonstrates how to structure a Go application with repositories, services, and controllers.

## Features

- User CRUD operations (Create, Read, Update, Delete)
- API with clean architecture (Controller, Service, Repository)
- Simple database connection with SQLite
- Request validation

## Table of Contents

1. [Getting Started](#getting-started)
2. [Install Dependencies](#install-dependencies)
3. [Run the Application](#run-the-application)
4. [API Endpoints](#api-endpoints)
5. [Testing](#testing)
6. [Contributing](#contributing)
7. [License](#license)

## Getting Started

Follow the steps below to set up this project locally.

### 1. Clone the Repository

```bash
git clone https://github.com/fatihbahadir/go-user-crud.git
cd go-user-crud
```

### 2. Install Dependencies

Before running the application, you need to install the required Go dependencies.

```bash
go mod tidy
```

This will install all the dependencies listed in `go.mod` and `go.sum`.

### 3. Set Up the Database

The application uses SQLite as the database by default. Make sure you have SQLite installed on your system. If you're using a different database, you may need to modify the configuration file (`config/database.go`) to use your preferred database (e.g., PostgreSQL, MySQL).

### 4. Run the Application

To start the application, run the following command:

```bash
go run main.go
```

This will start the server on `http://localhost:8888`.

## API Endpoints

### 1. Create User

- **POST** `/user`
  
Create a new user with the required fields: `name`, `surname`, `email`, `phone_number`.

#### Request Body Example:

```json
{
  "name": "John",
  "surname": "Doe",
  "email": "john.doe@example.com",
  "phone_number": "1234567890"
}
```

#### Response:

```json
{
    "code": 201,
    "message": "User created successfully"
}
```

### 2. Get All Users

- **GET** `/user`

Get a list of all users.

#### Response Example:

```json
[
  "code": 200,
  "message": "Users fetched successfully",
  "data" : [
  {
    "id": "550e8400-e29b-41d4-a716-446655440002",
    "name": "John",
    "surname": "Doe",
    "email": "john.doe@example.com",
    "phone_number": "05111111111",
    "created_at": "2022-01-01T00:00:00Z"
  },
  {
    "id": "550e8400-e29b-41d4-a716-446655440001",
    "name": "Elif",
    "surname": "Çelik",
    "email": "elifcelik@gmail.com",
    "phone_number": "05121121212",
    "created_at": "2024-12-31T20:38:40Z"
  }
  ]        
]
```

### 3. Get User by ID

- **GET** `/user/{id}`

Get a user by their ID.

#### Response Example:

```json
{
  "code": 200,
  "message": "User found successfully",
  "data" : {
      "id": "uuid",
      "name": "John",
      "surname": "Doe",
      "email": "john.doe@example.com",
      "phone_number": "1234567890",
      "created_at": "2022-01-01T00:00:00Z"
  }
}
```

### 4. Update User

- **PATCH** `/user/{id}`

Update the user information. At least one field should be provided in the request body.

#### Request Body Example:

```json
{
  "name": "John Updated",
  "email": "john.updated@example.com"
}
```

#### Response:

```json
{
  "code": 200,
  "message": "User updated successfully",
  "data" : {
      "id": "uuid",
      "name": "John Updated",
      "surname": "Doe",
      "email": "john.updated@example.com",
      "phone_number": "1234567890",
      "created_at": "2022-01-01T00:00:00Z"
  }
}
```

### 5. Delete User

- **DELETE** `/user/{id}`

Delete a user by their ID.

#### Response:

```json
{
  "code": 200,
  "message": "User deleted successfully"
}
```

## Testing

To run tests, use the following command:

```bash
go test ./...
```

This will run all the tests in the repository.

## Contributing

Feel free to fork this project and submit pull requests. Contributions are always welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.