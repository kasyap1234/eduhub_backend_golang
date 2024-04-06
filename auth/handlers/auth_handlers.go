package auth

import (
    "github.com/supertokens/supertokens-golang/supertokens"
    "github.com/supertokens/supertokens-golang/recipe/emailpassword"
    "github.com/supertokens/supertokens-golang/recipe/session"
    "github.com/gin-gonic/gin"
    "os"
)

// InitSuperTokens initializes SuperTokens with the provided configuration
func InitSuperTokens(router *gin.Engine) error {
    // Initialize SuperTokens
    err := supertokens.Init(supertokens.TypeInput{
        Supertokens: &supertokens.ConnectionInfo{
            ConnectionURI: "https://try.supertokens.com", // Replace with your SuperTokens connection URI
        },
        APPInfo: supertokens.AppInfo{
            AppName:        os.Getenv("APP_NAME"),
            APIDomain:      os.Getenv("API_DOMAIN"),
            WebsiteDomain:  os.Getenv("WEBSITE_DOMAIN"),
            APIBasePath:    "/auth",
            WebsiteBasePath: "/auth",
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
            c.Next()
        })).ServeHTTP(c.Writer, c.Request)
        c.Abort()
    })

    return nil
}
