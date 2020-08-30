init:
	go run github.com/99designs/gqlgen init

generate:
	go run github.com/99designs/gqlgen && go run ./plugins/plugins.go
