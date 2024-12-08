# syntax=docker/dockerfile:1
FROM postgres:alpine
COPY docker/dataset/init.sql /docker-entrypoint-initdb.d/
# COPY docker/dataset/seed.sql /docker-entrypoint-initdb.d/
