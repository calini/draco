package main

import (
	"github.com/calini/draco/setup"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	version = "v1"
	addr    = ":8080"
)

func main() {
	r := setup.Router()

	api := r.Group("/api/" + version)
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "peawng"})
		})
	}

	if err := setup.RunDefault(r, addr); err != nil {
		log.Fatal(err)
	}
}
