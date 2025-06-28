.PHONY: build run up

build:
	go build -o build/main cmd/main.go

run: build
	./build/main

up:
	air -c .air.toml


.PHONY: get_health get_me
get_health:
	curl -i -X GET http://localhost:8080/api/v1/health

get_me:
	curl -i -X GET http://localhost:8080/api/v1/users/me    

get_profile_me:
	curl -i -X GET http://localhost:8080/api/v1/users/profile/me

get_product_by_id:
	curl -i -X GET http://localhost:8080/api/v1/products/10010000
