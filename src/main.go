package main

import (
	"github.com/mateusVedoy/go-my-songs.git/src/application/controller"
	"github.com/mateusVedoy/go-my-songs.git/src/application/converter"
	"github.com/mateusVedoy/go-my-songs.git/src/application/router"
	usecase "github.com/mateusVedoy/go-my-songs.git/src/application/useCase"
)

func main() {
	route := inject()
	route.Start()
}

func inject() *router.Route {
	converter := converter.NewConverterImpl()
	manager := usecase.NewManager(converter)
	controller := controller.NewController(manager)
	route := router.NewRoute(controller)

	return route
}
