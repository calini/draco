package main

import (
	"github.com/calini/draco/integrations"
	"github.com/calini/draco/setup"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
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

		// TODO swtich to actual calls
		api.StaticFile("/general", "json/general.json")

		// Snyk Integration
		snykToken, ok := os.LookupEnv("SNYK_TOKEN")
		if !ok {
			log.Warnf("env var SNYK_TOKEN required for Snyk Integration")
		}
		api.Group("/snyk").GET("/", integrations.ProxySnyk(snykToken))
}

	if err := setup.RunDefault(r, addr); err != nil {
		log.Fatal(err)
	}
}
