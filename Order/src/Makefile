

all: clean

clean: 
	find . -name 'backend' -type f -exec rm -f {} \; 
	go clean

run:
	go run src/backend/$(app).go

main:
	go run src/backend/main.go

format:
	go fmt backend

install:
	go install backend

go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v github.com/streadway/amqp
	go get -v go.mongodb.org/mongo-driver/bson
	go get -v go.mongodb.org/mongo-driver/bson/primitive
	go get -v go.mongodb.org/mongo-driver/mongo
	go get -v github.com/rs/cors

build:
	go build backend

start:
	./backend 
