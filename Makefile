backend:
	go run main.go
migrate-up:
	migrate -path Migrations -database "postgres://postgres:helloworld@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path Migrations -database "postgres://postgres:helloworld@localhost:5432/postgres?sslmode=disable" down

migrate-up-force:
	migrate -path Migrations -database "postgres://postgres:helloworld@localhost:5432/postgres?sslmode=disable" -verbose up 

migrate-down-force:
	migrate -path Migrations -database "postgres://postgres:helloworld@localhost:5432/postgres?sslmode=disable" -verbose down 

migrate-fix:
	migrate -path Migrations -database "postgres://postgres:helloworld@localhost:5432/postgres?sslmode=disable" force 1 -y

docker-build:
	docker build -t backend .

DB-up:
	docker-compose -f ./docker-compose-DB-Only.yaml up

DB-down:
	docker-compose -f ./docker-compose-DB-Only.yaml down

.PHONY: backend migrate-up migrate-down docker-build DB-up DB-down