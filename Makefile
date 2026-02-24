migrate:
	migrate -path db/migrations -database "postgres://root:mysecretpassword@localhost:5432/aj_market?sslmode=disable" up
sqlc:
	sqlc generate