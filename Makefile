VERSION = 1.0.0

run_server:
	@echo "Start server"
	go run ./cmd/server/main.go

migrate:
	 migrate -source file://db/migrations -database "postgres://user:secret@localhost:5432/risk-assessor?sslmode=disable" up

#create migration
create_migration:
	migrate create -ext sql -dir db/migrations -seq create_users_table