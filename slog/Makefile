.PHONY: build run

build:
	go build -o build/main cmd/main.go

run: build
	./build/main


.PHONY: get_health get_users
get_health:
	curl -sX GET http://localhost:8080/api/v1/health | jq .

get_users:
	curl -sX GET http://localhost:8080/api/v1/users | jq .
