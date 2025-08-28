package main

import (
	"fmt"

	"github.com/ggualbertosouza/rate-limiter/src/internal/config"
)

func main() {
	config := config.LoadEnvConfig()

	fmt.Printf("%s %s", config.ServerPort, config.ServerHost)
}
