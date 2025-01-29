package main

import (
	"github.com/joho/godotenv"
	"log"
	"manulatorre98/trading/customErrors"
	"manulatorre98/trading/graph"
	"manulatorre98/trading/graph/directives"
	"manulatorre98/trading/repository"
	"manulatorre98/trading/repository/dolarRepository"
	"manulatorre98/trading/repository/userRepository"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func init() {
	customErrors.DefaultTranslation()
}

func main() {
	var err error
	err = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize database
	db, err := repository.PostgresDbInit()
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	var userRepo userRepository.UserRepository = userRepository.NewUserPsqlRepository(db)
	var dolarRepo dolarRepository.DolarRepository = dolarRepository.NewDolarPsqlRepository(db)

	c := graph.Config{Resolvers: &graph.Resolver{UserRepository: userRepo, DolarRepository: dolarRepo}}
	c.Directives.Binding = directives.Binding
	srv := handler.New(graph.NewExecutableSchema(c))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
