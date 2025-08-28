package main

import (
	"github.com/ggualbertosouza/rate-limiter/src/internal/config"
	"github.com/ggualbertosouza/rate-limiter/src/internal/server"
)

func main() {
	config.Init()
	server.InitServer()
}
