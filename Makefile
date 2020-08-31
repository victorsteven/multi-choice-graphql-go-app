init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen && go run ./app/models/model_tags/model_tags.go

run:
	go run main.go

test:
	go test -v ./...