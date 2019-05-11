package main

import (
	"github.com/LIYINGZHEN/gotemplate/configs"
	"github.com/LIYINGZHEN/gotemplate/internal/app/http"
)

func main() {
	c := configs.New()
	server := http.AppServer{}
	server.Run(c.Port)
}
