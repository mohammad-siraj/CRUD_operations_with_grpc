#!bin/sh
docker cp SQL_files/car.sql crud_grpc:/car.sql
docker exec -u root crud_grpc psql root root -f car.sql