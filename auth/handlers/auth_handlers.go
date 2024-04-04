package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/recipe/password"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)
func main(){
	apiBasePath := "/auth"
	websiteBasePath :="/auth"
	err :=supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// replace try.supertokens.com with self hosted one ;
			ConnectionURI: "https://try.supertokens.com",

		},
		APPInfo: supertokns.AppInfo{
			err:= godotenv.Load()
			if err!=nil{
				log.Fatal("Error loading .env file")
			}
			AppName: os.Getenv("APP_NAME")
			APIDomain: os.Getenv("API_DOMAIN")
			WebsiteDomain: os.Getenv("WEBSITE_DOMAIN")
			APIBasePath: &apiBasePath,
			WebsiteBase
		}
	})
}
