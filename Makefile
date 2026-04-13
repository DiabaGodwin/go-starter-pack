DB_URL=postgres://postgres:postgres123@localhost:5432/mydb?sslmode=disable

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" down 1

migrate-force:
	migrate -path db/migrations -database "$(DB_URL)" force $(version)