

# Catalogue Service

A simple microservice with endpoints to get/add information about, add new, and delete items off a database.

This application is built using **GoLang**, **Chi** router for handling HTTP requests, and **GORM** with **PostgreSQL** for managing the database.

## Endpoints:
- **Create**: Create a new item, with initalized values for each field (including an indexed uuid).
- **Get**: Retrieve the feilds of an item through its ID.
- **Update**: Update the value of any number of feilds of a particular item, through its ID .
- **Delete**: Permanatly delete an item from the database using its ID.

## How to Run the Application

### Prerequisites
- **GoLang** (version 1.18 or higher)
- **PostgreSQL** (or a Dockerized PostgreSQL instance)

### 1. Clone the repository

```sh
git clone https://github.com/kalabyebeltal/catalogue-api.git
cd catalogue-api
```

### 2. Set up environment variables

Create a `.env` file in the root of the project to store your database credentials and other configuration settings. Check `.env.example` file

### 3. Install Go dependencies

Make sure Go modules are initialized and dependencies are installed:

```sh
go mod tidy
```

### 4. Run the application locally

To run the application locally:

```sh
go run cmd/main.go
```

The application will start on the default port **8080**. You can visit the following endpoints:

- `http://localhost:8080/catalogue`
- `http://localhost:8080/catalogue/post`
- `http://localhost:8080/catalogue/get{id}`
- `http://localhost:8080/catalogue/update{id}`
- `http://localhost:8080/catalogue/delete{id}`
