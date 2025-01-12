package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"manulatorre98/trading/dependencyContainer"
	Middlewares "manulatorre98/trading/middlewares"
	"manulatorre98/trading/repository"
)

//var db *sql.DB

func main() {
	// Load .env file
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize gin router
	router := gin.Default()
	router.Use(Middlewares.ErrorHandlingMiddleware)

	// Initialize database
	db, err := repository.PostgresDbInit()
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Initialize dependency container
	container := dependencyContainer.NewDependencyContainer(db)

	// Routes
	router.POST("/users", container.UserCtrl.InsertUser)
	router.GET("/users", container.UserCtrl.GetUser)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal("Error starting server")
	}

}
