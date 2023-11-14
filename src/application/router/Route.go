package router

import (
	"fmt"
	"net/http"

	"github.com/mateusVedoy/go-my-songs.git/src/application/controller"
)

type Route struct {
	controller controller.Controller
}

func (r Route) Start() {
	http.HandleFunc("/music/add", r.controller.Add)
	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func NewRoute(controller *controller.Controller) *Route {
	return &Route{controller: *controller}
}
