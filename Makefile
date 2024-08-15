.PHONY: run-docker
run-docker:
	docker compose up --force-recreate --no-deps --build

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o image_shifter

.PHONY: start
start:
	./image_shifter
