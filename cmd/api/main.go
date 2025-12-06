package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tdalexm/goson-server/internal/repository"
	"github.com/tdalexm/goson-server/internal/services"
)

func main() {
	router := gin.Default()
	jsonRepo := repository.NewJsonRepo("./data/db.json")

	data, _ := jsonRepo.Load()
	stateRepo := repository.NewStateRepository(data)

	handler := &Handler{
		listSR:       *services.NewListService(stateRepo),
		listFilterSR: *services.NewListFilterService(stateRepo),
		getSR:        *services.NewGetService(stateRepo),
	}

	router.GET("/:resource", handler.List)
	router.GET("/:resource/:id", handler.Get)

	log.Fatalln(router.Run())
}
