# Go Fiber Notes API

A secure and minimal REST API built with Go Fiber, MySQL, and JWT for personal note management.

## Features

- User regusstration and login with hashed passwords (bcrypt)
- JWT-based authentication to protect routes
- Create, read, update, and delete personal notes
- Only logged-in users can access their own notes
- Pagination and search support for notes
- Docker support for easy deployment
- CLI tool to seed sample data

## Getting Started

### Run Locally

1. Create a `.env` file in the root directory (you can copy from `.env.example`)
2. Make sure MySQL is running and your DB credentials are correct
3. Install dependencies:
   go mod tidy
4. Run the app:
   go run main.go

### Run with Docker

1. Make sure `.env` exists
2. Build and start:
   docker-compose up --build

### Seed the Database

To add sample users and notes for testing:
   go run cli/seed.go

## API Endpoints

POST /register – Register a new user  
POST /login – Log in and get JWT token  
POST /notes – Create a new note (requires JWT)  
GET /notes – Get all personal notes (with pagination and search)  
GET /notes/:id – Get a specific note by ID  
PUT /notes/:id – Update a note  
DELETE /notes/:id – Delete a note

## License

MIT

# # Go Fiber Notes API

A secure and minimal REST API built with Go Fiber, MySQL, and JWT for personal note management.

## Features

- User regusstration and login with hashed passwords (bcrypt)
- JWT-based authentication to protect routes
- Create, read, update, and delete personal notes
- Only logged-in users can access their own notes
- Pagination and search support for notes
- Docker support for easy deployment
- CLI tool to seed sample data

## Getting Started

### Run Locally

1. Create a `.env` file in the root directory (you can copy from `.env.example`)
2. Make sure MySQL is running and your DB credentials are correct
3. Install dependencies:
   go mod tidy
4. Run the app:
   go run main.go

### Run with Docker

1. Make sure `.env` exists
2. Build and start:
   docker-compose up --build

### Seed the Database

To add sample users and notes for testing:
   go run cli/seed.go

## API Endpoints

POST /register – Register a new user  
POST /login – Log in and get JWT token  
POST /notes – Create a new note (requires JWT)  
GET /notes – Get all personal notes (with pagination and search)  
GET /notes/:id – Get a specific note by ID  
PUT /notes/:id – Update a note  
DELETE /notes/:id – Delete a note

## License

MIT
