all: wigo

wigo:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o wigo wigo.go

wigo-worker:
	CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o wigo-worker worker/main.go

start-app:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	# Install godotenv with 'go install github.com/joho/godotenv/cmd/godotenv@latest'
	reflex -s -r '\.go$$' -- godotenv -f .env go run wigo.go

start-view:
	# Install reflex with 'go install github.com/cespare/reflex@latest'
	reflex -r '\.qtpl$$' -- qtc -dir=internal/view

db-migrate:
	migrate -path migrations -database "postgres://127.0.0.1/wigo?sslmode=disable" up

db-schema-dump:
	pg_dump --schema-only -O wigo > internal/database/schema.sql

sqlc-gen:
	sqlc --experimental generate

.PHONY: wigo wigo-worker start-app start-view db-migrate db-schema-dump sqlc-gen
