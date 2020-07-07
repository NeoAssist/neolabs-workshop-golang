package server

import (
	"fmt"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/database/store"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/server/handler"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/pkg/server/router"
	"github.com/NeoAssist/neolabs-workshop-golang/internal/utils"
)

// RunServer Initialize a new server
func RunServer() {
	r := router.InitializeRoutes()
	v1 := r.Group("/api/v1")

	d := database.New()
	database.AutoMigrate(d)

	posterStore := store.NewPosterStore(d)
	h := handler.NewHandler(posterStore)
	h.Register(v1)

	address := fmt.Sprintf("%v:%v", utils.GetEnv("API_HOST"), utils.GetEnv("API_PORT"))

	r.Logger.Fatal(r.Start(address))
}
