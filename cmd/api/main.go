package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/repository"
	"github.com/tdalexm/goson-server/internal/services"
)

func main() {
	// Comment for debug mode
	gin.SetMode(gin.ReleaseMode)

	dbPath := flag.String("db", "db.json", "Path to JSON database file")
	port := flag.String("port", "8080", "Port to run the server")
	help := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *help {
		fmt.Println("goson-server - A JSON server implementation in Go")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	log.Printf("Starting server with db: %s on port: %s", *dbPath, *port)

	if _, err := os.Stat(*dbPath); os.IsNotExist(err) {
		log.Fatalf("Database file not found: %s", *dbPath)
	}

	router := gin.Default()
	jsonRepo := repository.NewJsonRepo(*dbPath)

	data, _ := jsonRepo.Load()
	stateRepo := repository.NewStateRepository(data)

	handler := &Handler{
		listSR:        *services.NewListService(stateRepo),
		listFilterSR:  *services.NewListFilterService(stateRepo),
		getSR:         *services.NewGetService(stateRepo),
		createSR:      *services.NewCreateService(stateRepo),
		updateSR:      *services.NewUpdateService(stateRepo),
		updateFieldSR: *services.NewUpdateFieldsService(stateRepo),
		deleteSR:      *services.NewDeleteService(stateRepo),
	}

	router.GET("/:resource", handler.List)
	router.GET("/:resource/:id", handler.Get)
	router.POST("/:resource", handler.Create)
	router.POST("/:resource/:id", handler.Update)
	router.PATCH("/:resource/:id", handler.Update)
	router.DELETE("/:resource/:id", handler.Delete)

	log.Fatalln(router.Run(":" + *port))
}
