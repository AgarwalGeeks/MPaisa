.PHONY: createdb dropdb sqlc

createdb: 
	docker exec -it postgres_server createdb --username=haagarwa --owner=haagarwa MPaisa

dropdb:
	docker exec -it postgres_server dropdb --username=haagarwa MPaisa

postgres:
	docker run --name postgres_server -e POSTGRES_USER=haagarwa -e POSTGRES_PASSWORD=Harshit@12345 -p 5435:5432 -d postgres:latest

migrateupinit:
	migrate -path db/migration -database "postgresql://haagarwa:Harshit@12345@localhost:5435/MPaisa?sslmode=disable" -verbose up:

migratedowninit:
	migrate -path db/migration -database "postgresql://haagarwa:Harshit@12345@localhost:5435/MPaisa?sslmode=disable" -verbose down

sqlc:
	sqlc generate