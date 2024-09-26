POSTGRES_CONNECTION_STRING := postgresql://root:secret@localhost:5432/simple_go?sslmode=disable

init:
	docker run --name simple_go -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d -p 5432:5432  postgres:16-alpine
	until docker exec -it simple_go pg_isready; do sleep 1; done
	sleep 1
	docker exec -it simple_go createdb -U root simple_go

clean:
	docker stop simple_go
	docker rm simple_go

start:
	docker start simple_go
	make run
	
stop:
	docker stop simple_go

mg-file:
	migrate create -ext sql -dir database/migrations -seq $(name)

migrate-up:
	migrate -path database/migrations -database $(POSTGRES_CONNECTION_STRING) -verbose up

migrate-down:
	migrate -path database/migrations -database $(POSTGRES_CONNECTION_STRING) -verbose down

sqlc:
	sqlc generate

run:
	go run main.go

test:
	go test -v -cover ./...

.PHONY: mg-file migrate-up migrate-down sqlc init clean start stop run test