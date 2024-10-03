## PayFam Project
This is a Go project designed to handle a web service that connects to a PostgreSQL database. The project uses GORM (Go's ORM library) for database operations and can be run inside a Docker container for easy setup and portability.

<img width="1435" alt="Screenshot 2024-10-03 at 1 48 02 PM" src="https://github.com/user-attachments/assets/8ae03723-3259-451b-9ba9-00c4cf939d55">
<img width="1435" alt="Screenshot 2024-10-03 at 1 48 14 PM" src="https://github.com/user-attachments/assets/638a63d9-9e05-4260-b9b8-b879a3bbbffb">
<img width="1205" alt="Screenshot 2024-10-03 at 1 50 51 PM" src="https://github.com/user-attachments/assets/b5cc5c39-01a9-48b6-ab25-ad66a29fbefb">
<img width="1205" alt="Screenshot 2024-10-03 at 1 51 22 PM" src="https://github.com/user-attachments/assets/b130683a-624f-4ae4-bf58-bb0a127bec11">

## Features
Golang for backend logic
PostgreSQL as the database
GORM as the ORM for database interactions
Docker for containerization

## Prerequisites
Before you begin, ensure you have met the following requirements:

Go (Golang) installed: Download Go
Docker installed: Install Docker
PostgreSQL installed locally or running in a Docker container
## Project Structure

├── Dockerfile          # Docker configuration to containerize the Go app
├── docker-compose.yml  # Docker Compose file to set up the app and PostgreSQL
├── go.mod              # Go module file
├── go.sum              # Dependency lock file
├── main.go             # Main entry point for the Go application
├── internal/
│   ├── entity/
│   │   └── videoEntity.go       # Database connection and initialization logic
│   └── dao/
│       └── videoDao.go    # Video model and data operations (GORM)
└── README.md           # Project documentation (this file)

## Getting Started
Clone the Repository
git clone https://github.com/yourusername/yourproject.git
cd yourproject

## Set Up Environment Variables
Create a .env file in the root of your project and add the following configuration. Adjust the values according to your setup:

env
DATABASE_URL=postgres://user:password@localhost:5432/mydatabase?sslmode=disable
DB_HOST=localhost
DB_USER=user
DB_PASSWORD=password
DB_NAME=mydatabase
DB_PORT=5432
Build and Run with Docker
If you are using Docker, follow these steps:

## Build the Docker image:

docker build -t my-go-app .
Run the Docker container: You can run the app and PostgreSQL together using Docker Compose:

docker-compose up --build
This will start the Go application and PostgreSQL database in containers.

Running Locally (Without Docker)
If you prefer to run the project on your local machine without Docker:

## Install Go dependencies:

go mod tidy
## Run the Go application:

go run main.go
Ensure that PostgreSQL is running locally and accessible at the DATABASE_URL defined in your environment variables.

## Database Migrations
This project uses GORM for handling database migrations. Once the app is running, GORM will automatically apply migrations for the Video model when connecting to the database.

To manually trigger migrations, ensure the InitDatabase function in internal/database/db.go calls:

db.AutoMigrate(&dao.Video{})
API Endpoints
You can add documentation here for the API endpoints your service exposes, for example:

GET /videos
GET /videos/view
GET /videos/search
