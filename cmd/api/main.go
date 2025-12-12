package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	jsonloader "github.com/tdalexm/goson-server/internal/adapters/driven/json_loader"
	"github.com/tdalexm/goson-server/internal/adapters/driven/repository"
	driverhttp "github.com/tdalexm/goson-server/internal/adapters/driver/http"
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
	jsonRepo := jsonloader.NewJsonRepo(*dbPath)

	data, _ := jsonRepo.Load()
	stateRepo := repository.NewStateRepository(data)

	handler := &driverhttp.Handler{
		ListSR:        *services.NewListService(stateRepo),
		ListFilterSR:  *services.NewListFilterService(stateRepo),
		GetSR:         *services.NewGetService(stateRepo),
		CreateSR:      *services.NewCreateService(stateRepo),
		UpdateSR:      *services.NewUpdateService(stateRepo),
		UpdateFieldSR: *services.NewUpdateFieldsService(stateRepo),
		DeleteSR:      *services.NewDeleteService(stateRepo),
	}

	router.GET("/:collection", handler.List)
	router.GET("/:collection/:id", handler.Get)
	router.POST("/:collection", handler.Create)
	router.POST("/:collection/:id", handler.Update)
	router.PATCH("/:collection/:id", handler.Update)
	router.DELETE("/:collection/:id", handler.Delete)

	log.Fatalln(router.Run(":" + *port))
}
