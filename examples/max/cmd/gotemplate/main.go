package main

import (
	"github.com/go-project-action/simple-webpage/examples/max/configs"
	"github.com/go-project-action/simple-webpage/examples/max/internal/app/http"
)

func main() {
	c := configs.New()
	server := http.AppServer{}
	server.Run(c.Port)
}
