DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

sqlcwin:
	sqlc generate

sqlcdocker:
	 docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

server:
	go run main.go

test:
	go test -v -cover -short ./...

mock:
	mockgen -package mockdb -destination db/mock/Store.go simpleBank/db/sqlc Store

.PHONY: migrateup migratedown sqlcwin sqlcdocker server test mock