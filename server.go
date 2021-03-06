package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mojeico/gqlgen-golang/graph"
	"github.com/mojeico/gqlgen-golang/graph/generated"
	customMiddleware "github.com/mojeico/gqlgen-golang/internal/middleware"
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

	router := chi.NewRouter()
	/*	router.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:8080"},
		OptionsPassthrough: true,
		//Debug:              true,
	}).Handler)*/

	//router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Group(func(router chi.Router) {
		router.Use(customMiddleware.AuthMiddleware(userRepo))
		router.Handle("/query", srv)
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
