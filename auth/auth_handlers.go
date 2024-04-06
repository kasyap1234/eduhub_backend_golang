package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

// InitSuperTokens initializes SuperTokens with the provided configuration
func InitSuperTokens(router *gin.Engine, apiBasePath, websiteBasePath string) error {
	// Initialize SuperTokens
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "https://try.supertokens.com", // Replace with your SuperTokens connection URI
		},
		AppInfo: supertokens.AppInfo{
			AppName:         os.Getenv("APP_NAME"),
			APIDomain:       os.Getenv("API_DOMAIN"),
			WebsiteDomain:   os.Getenv("WEBSITE_DOMAIN"),
			APIBasePath:     &apiBasePath,     // Use pointer to string
			WebsiteBasePath: &websiteBasePath, // Use pointer to string
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
		},
	})
	if err != nil {
		return err
	}

	// Attach SuperTokens middleware to the router
	router.Use(func(c *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			// Your middleware logic here
		})).ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	return nil
}
