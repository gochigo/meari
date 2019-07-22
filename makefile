

run:
	go run cmd/httpd/main.go

build:
	go build ./cmd/httpd

image:
	sudo docker build -t meari:latest .
