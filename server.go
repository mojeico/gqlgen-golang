package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mojeico/gqlgen-golang/graph"
	"github.com/mojeico/gqlgen-golang/graph/generated"
	"github.com/mojeico/gqlgen-golang/internal/repository"
	"github.com/mojeico/gqlgen-golang/internal/service"
	"github.com/mojeico/gqlgen-golang/pkg/database"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

var (
	mongo = database.NewMongo()

	userRepo    = repository.NewUserRepo(mongo)
	userService = service.NewUserService(userRepo)

	meetupRepo    = repository.NewMeetupsRepo(mongo)
	meetupService = service.NewMeetupsRepo(meetupRepo)
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var srv = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:    userService,
		MeetupsService: meetupService,
	},
	}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
