// auth: kunlun
// date: 2019-01-15
// description:
package middle

import (
	"encoding/base64"
	"net/http"
	"strings"
	"yihuo"
)

func Auth(context *yihuo.Context, res http.ResponseWriter, req *http.Request) (bool, string) {
	c := context.StrategyInfo
	if strings.ToLower(c.Auth) == "basic" {
		authStr := []byte(c.BasicUserName + ":" + c.BasicPassword)
		authorization := "Basic " + base64.StdEncoding.EncodeToString(authStr)
		auth := strings.Join(req.Header["Authorization"], ", ")
		if authorization != auth {
			return false, "Username or UserPassword Error"
		}
	} else if strings.ToLower(c.Auth) == "apikey" {
		apiKey := strings.Join(req.Header["Apikey"], ", ")
		if c.ApiKey != apiKey {
			return false, "Invalid ApiKey"
		}
	}
	return true, ""
}
