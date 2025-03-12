CREATE DATABASE etl_db;

\c etl_db;

CREATE TABLE IF NOT EXISTS raw_data (
    id SERIAL PRIMARY KEY,
    data JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transformed_data (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255),
    dob TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);
