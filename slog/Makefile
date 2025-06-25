.PHONY: build run

build:
	go build -o build/main cmd/main.go

run: build
	./build/main


.PHONY: get_health get_me
get_health:
	curl -i -X GET http://localhost:8080/api/v1/health

get_me:
	curl -i -X GET http://localhost:8080/api/v1/users/me    
