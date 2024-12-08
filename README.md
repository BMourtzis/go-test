#go-test

## Docker commands
- Create Docker Containers: docker-compose up --force-recreate --build -d
- Create Docker Containers with debugger: docker-compose -f ./docker-compose.debug.yaml up --force-recreate --build -d
- Open console on Container: docker exec -it <container_name or container_id> <command>
    - docker exec -it a9ef189b9bcbe77453e28964542e0ed3d435eb86218809d75ba4634cf3aa99e2 psql -U postgres -W services_db

## Postgres commands
- List all tables: SELECT * FROM pg_catalog.pg_tables;