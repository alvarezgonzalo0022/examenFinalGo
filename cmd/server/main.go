package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	routes "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/router"
	_ "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/docs"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/middleware"
	"github.com/alvarezgonzalo0022/examenFinalGo/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	puerto = "8080"
)

// @title Dental Clinic Lisa Necesita Frenos
// @version 1.0
// @description This is a sample dental clinic API for managing appointments, dentists, and patients.
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// Recover from panic.
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	// Load the environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database.
	db := store.ConnectDB()

	// Create a new Gin engine.
	router := gin.New()
	router.Use(gin.Recovery())
	// Add the logger middleware.
	router.Use(middleware.Logger())

	// Add the swagger handler.
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the application.
	runApp(db, router)

	// Close the connection.
	defer db.Close()

}

func runApp(db *sql.DB, engine *gin.Engine) {
	// Run the application.
	router := routes.NewRouter(engine, db)
	// Map all routes.
	router.MapRoutes()
	if err := engine.Run(fmt.Sprintf(":%s", puerto)); err != nil {
		panic(err)
	}

}
