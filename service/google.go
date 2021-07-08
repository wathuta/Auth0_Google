package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "http://localhost:9090/callback-gl",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	oauthStateStringGl = ""
)

func InitializeOauthGoogle() {
	oauthConfGl.ClientID = viper.GetString("google.clientID")
	oauthConfGl.ClientSecret = viper.GetString("google.clientSecret")
	oauthStateStringGl = viper.GetString("oauthStateString")
}
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	HandleLogin(w, r, oauthConfGl, oauthStateStringGl)
}

/*
CallBackFromGoogle Function
*/
func CallBackFromGoogle(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")

	if state != oauthStateStringGl {
		log.Println("invalid oauth state, expected " + oauthStateStringGl + ", got " + state + "\n")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	//logger.Log.Info(code)

	if code == "" {
		log.Println("Code not found...")
		w.Write([]byte("Code Not Found to provide AccessToken..\n"))
		reason := r.FormValue("error_reason")
		if reason == "user_denied" {
			w.Write([]byte("user denied permission"))
		}
	} else {
		token, err := oauthConfGl.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Println("oauthConfGl.Exchange() failed with " + err.Error() + "\n")
			return
		}
		/*	logger.Log.Info("TOKEN>> AccessToken>> " + token.AccessToken)
			logger.Log.Info("TOKEN>> Expiration Time>> " + token.Expiry.String())
			logger.Log.Info("TOKEN>> RefreshToken>> " + token.RefreshToken)
		*/resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
		if err != nil {
			log.Println("Get: " + err.Error() + "\n")
		}
		defer resp.Body.Close()
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ReadAll: " + err.Error() + "\n")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		log.Println("parseResponseBody: " + string(response) + "\n")

		w.Write([]byte("Hello, I'm protected\n"))
		w.Write([]byte(string(response)))
		return
	}

}
