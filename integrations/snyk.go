package integrations

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	snykURL = "https://snyk.io/"
	snykAPIPath = "api/v1/"
)

func ProxySnyk(token string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u, err := url.Parse(snykURL)
		if err != nil {
			log.Info(err)
			c.String(http.StatusInternalServerError, "internal server error")
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(u)
		c.Request.Host = u.Host
		c.Request.URL.Path = snykAPIPath + c.Param("path")
		c.Request.Header["Authorization"] = []string{"token " + token}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}