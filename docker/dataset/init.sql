-- CREATE USER IF NOT EXISTS postgres;

-- CREATE DATABASE IF NOT EXISTS services_db;
GRANT ALL PRIVILEGES ON DATABASE services_db TO postgres;

\connect services_db;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(50),
    lastname VARCHAR(50),
    age INT
); 

CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    isActive BIT
); 

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO postgres;
