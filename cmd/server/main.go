package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	routes "github.com/alvarezgonzalo0022/examenFinalGo/cmd/server/router"
	// _ "github.com/alvarezgonzalo0022/examenFinalGo/docs"
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
