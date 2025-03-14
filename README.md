# Go ETL Pipeline 🚀

## 📌 Overview  
This project is a **Go-based ETL (Extract, Transform, Load) pipeline** that:  
✅ Fetches user data from an **external API (`https://randomuser.me/api/`)** every 30 seconds  
✅ Stores **raw JSON responses** in **PostgreSQL** and in local storage (`data/raw/`)  
✅ Uses **Go channels & goroutines** for **parallel data processing**  
✅ Transforms the data (extracts `name`, `email`, `dob`)  
✅ Saves **processed data in JSON format (`data/processed/`)**  
✅ Exposes a **health check endpoint** (`http://localhost:8080/health`)  

---

## 📂 Project Structure  
```plaintext
📂 etl-pipeline
 ├── 📂 config/                
 │    ├── config.go           # Loads database & API configurations
 ├── 📂 data/                    
 │    ├── raw/                # Contain raw data
 │    ├── processed/          # Contain processed data
 ├── 📂 db/                    
 │    ├── db.go               # Handles PostgreSQL connection
 │    ├── init.sql            # Creates PostgreSQL tables on startup
 ├── 📂 services/              
 │    ├── fetch.go            # Fetches raw API data
 │    ├── transform.go        # Transforms the data
 │    ├── store.go            # Saves processed data
 ├── 📂 utils/                 
 │    ├── helper.go           # Helper utilities
 ├── 📂 logs/                   # Contain log file  
 ├── main.go                   # Application entry point (goroutines & channels)
 ├── Dockerfile                # Docker containerization
 ├── docker-compose.yml        # Manages ETL & PostgreSQL containers
 ├── go.mod                    # Go dependencies
 ├── README.md                 # Documentation
```
## 📌 Prerequisites

Before running the application, make sure you have:

- Go 1.18+ → Install from golang.org
- Docker & Docker Compose → Install from docker.com
- PostgreSQL → Either install locally or use Docker

## Running the ETL Pipeline (Without Docker)

- Step 1: Configure Database
1️⃣ Start PostgreSQL and create a database:
```
psql -U postgres -c "CREATE DATABASE etl_db;"
```
1️⃣ Create the required table:
```
CREATE TABLE raw_data (
    id SERIAL PRIMARY KEY,
    data JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```
- Step 2: Setup Environment
Modify config/config.go to match your PostgreSQL credentials.

- Step 3: Run the Application
```
go run main.go
```
## Running with Docker
- Step 1: Build & Start the Containers
```
docker-compose up --build
```
This will:

✅ Start PostgreSQL and automatically create the required tables

✅ Start the Go ETL app, fetching and processing data every 30 seconds

- Step 2: Monitoring the ETL Pipeline

1. Check Running Containers
```
docker ps
```

2. Check Processed Data
```
cat data/processed/processed_data.json
```

3. Check service health
```
curl http://localhost:8080/health
```