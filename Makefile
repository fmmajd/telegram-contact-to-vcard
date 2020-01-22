all: say_hello

say_hello:
	@echo "Hello World!"

up:
	@docker-compose up -d

bash:
	@docker-compose exec go bash

run:
	@docker-compose exec go bash -c "go run main.go"

build:
	@docker-compose exec go bash -c "go build main.go"
