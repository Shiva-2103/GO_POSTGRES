# GO_POSTGRES

A Go-based REST API application that integrates with PostgreSQL to manage stocks. This project demonstrates a simple CRUD (Create, Read, Update, Delete) operation setup using Go, Gorilla Mux, and PostgreSQL.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)


---

## Features

- Manage stocks with basic CRUD operations.
- RESTful API structure.
- Built with:
  - [Go](https://golang.org/)
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [PostgreSQL](https://www.postgresql.org/)

---

## Prerequisites

1. [Go (1.20+)](https://golang.org/doc/install)
2. [PostgreSQL](https://www.postgresql.org/download/)
3. Basic knowledge of REST APIs and Go programming.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Shiva-2103/GO_POSTGRES.git
   cd GO_POSTGRES
2. Install dependencies:
    ```bash
    go mod tidy
3. Set up PostgreSQL:

Create a database named stocksdb.
Run the following SQL script to create the required table:

    CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INT NOT NULL,
    company TEXT NOT NULL );

4. Configure the database connection in the project:
   Edit the database connection parameters in the middleware/db.go file.

## Usage

1. Run the application:
      ```bash
    go run main.go

2. The server will start at http://localhost:3000.

## API Endpoints
Base URL: http://localhost:3000
Method	Endpoint	Description
GET	/api/stock	Get all stocks
GET	/api/stock/{id}	Get stock by ID
POST	/api/newstock	Add a new stock
PUT	/api/stock/{id}	Update stock by ID
DELETE	/api/deletestock/{id}	Delete stock by ID

Example Payload for /api/newstock:
```
{
    "name": "TSLA",
    "price": 1000,
    "company": "Tesla"
}
```

## Project Structure
```
GO_POSTGRES/
├── main.go                # Entry point of the application
├── go.mod                 # Go module file
├── middleware/            # Business logic and database handling
│   ├── db.go              # Database connection
│   ├── handlers.go        # Handlers for API requests
├── router/                # Routes for the application
│   ├── router.go          # Router setup
└── README.md              # Project documentation
```


