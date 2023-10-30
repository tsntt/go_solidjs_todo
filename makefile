rundev:
	go run ./cmd/api/main.go

setuppostgres:
	docker start postgres14
	docker exec -it postgres14 createdb --username=postgres --owner=postgres todo_api

newmigration:
	migrate create -ext sql -dir data/postgres/migration -seq initial_tables

migrateup:
	migrate -path data/postgres/migration -database "postgresql://postgres:postgres@localhost:5432/todo_api?sslmode=disable" -verbose up

seedpostgres:
	go run cmd/seed/main.go -db=postgres