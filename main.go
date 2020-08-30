package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"multi-choice/app/interfaces"
	"multi-choice/app/domain/repository/answer"
	"multi-choice/app/domain/repository/question"
	"multi-choice/app/domain/repository/question_option"
	"multi-choice/app/generated"
	"multi-choice/app/infrastructure/db"
	"multi-choice/app/infrastructure/persistence"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}



func main() {

	var (
		defaultPort = "8080"
		databaseUser = os.Getenv("DATABASE_USER")
		databaseName = os.Getenv("DATABASE_NAME")
		databaseHost = os.Getenv("DATABASE_HOST")
		databasePort = os.Getenv("DATABASE_PORT")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//appAddr := ":" + os.Getenv("PORT")

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)

	conn := db.OpenDB(dbConn)

	var ansService answer.AnsService
	var questionService question.QuesService
	var questionOptService question_option.OptService

	ansService = persistence.NewAnswer(conn)
	questionService = persistence.NewQuestion(conn)
	questionOptService = persistence.NewQuestionOption(conn)


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		AnsService: ansService,
		QuestionService: questionService,
		QuestionOptionService: questionOptService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

