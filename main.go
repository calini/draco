package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/calini/draco/setup"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	oauth "golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	version = "v1"
	addr    = ":8080"
)

func main() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	githubClientId, ok := os.LookupEnv("GITHUB_CLIENT_ID")
	if !ok {
		log.Fatal("GITHUB_CLIENT_ID required")
	}
	githubClientSecret, ok := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !ok {
		log.Fatal("GITHUB_CLIENT_SECRET required")
	}

	// Some test stuff
	var (
		githubConfig = oauth.Config{
			ClientID:     githubClientId,
			ClientSecret: githubClientSecret,
			Endpoint:     github.Endpoint,
			RedirectURL:  "localhost:8080/oauth2/callback",
			Scopes:       []string{
				"user",
			},
		}
		state = "random" // TODO RANDOMIZE
	)


	r := setup.Router()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "static/index.html", nil)
	})

	loginApi := r.Group("/login")
	{
		// Redirect to Github
		loginApi.GET("/", func(c *gin.Context) {
			// Send user to GitHub with a randomised state to prevent CSRF
			url := githubConfig.AuthCodeURL(state)
			c.Redirect(http.StatusTemporaryRedirect, url)
		})

		// Redirect Back
		loginApi.GET("/callback", func(c *gin.Context) {

			// Check state match to prevent CSRF
			if c.PostForm("state") != state {
				log.Warn("state is not valid!")
				c.Redirect(http.StatusTemporaryRedirect, "")
				return
			}

			// Token exchange with GitHub
			token, err := githubConfig.Exchange(c, c.PostForm("code"))
			if err != nil {
				log.Warn("could not get token: %s\n", err.Error())
				c.Redirect(http.StatusTemporaryRedirect, "/")
				return
			}

			// Get personal info from GitHub
			http.Get(fmt.Sprintf("<GitHub's user info endpoint>?access_token=%s", token.AccessToken))
		})
	}

	if err := setup.RunDefault(r, addr); err != nil {
		log.Fatal(err)
	}
}
