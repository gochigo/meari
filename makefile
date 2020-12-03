
default:
	@echo "rules: run, build, and image"

run:
	go run cmd/meari/main.go

build:
	go build ./cmd/meari

image:
	sudo docker build -t meari:latest .
