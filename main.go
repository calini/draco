package main

import (
	"github.com/calini/draco/integrations"
	"github.com/calini/draco/setup"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const (
	version = "v1"
	addr    = ":8080"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	r := setup.Router()

	api := r.Group("/api/" + version)
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "peawng"})
		})

		// TODO swtich to actual calls
		api.StaticFile("/general", "json/general.json")
		api.StaticFile("/ports", "json/ip_scans.json")
		api.StaticFile("/ports/1", "json/ip_scan_detail.json")
		api.StaticFile("/vulns", "json/vuln_scans.json")
		api.StaticFile("/vulns/1", "json/vuln_scan_detail.json")

		// Snyk Integration
		snykToken, ok := os.LookupEnv("SNYK_TOKEN")
		if !ok {
			log.Warnf("env var SNYK_TOKEN required for Snyk Integration")
		}
		api.Group("/snyk").Any("*path", integrations.ProxySnyk(snykToken))
}

	if err := setup.RunDefault(r, addr); err != nil {
		log.Fatal(err)
	}
}
