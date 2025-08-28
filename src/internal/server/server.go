package server

import (
	"log"
	"net/http"

	"github.com/ggualbertosouza/rate-limiter/src/internal/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	})

	if err := router.Run(config.EnvConfig.ServerHost + config.EnvConfig.ServerPort); err != nil {
		log.Fatalf("Erro ao iniciar servidor")
	}
}
