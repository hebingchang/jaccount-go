package jaccount

import (
	"fmt"
	"golang.org/x/oauth2"
	"math/rand"
	"net/http"
	"os"
	"testing"
)

/**
* @Author: He Bingchang
* @Date: 2020/9/7
* @Description:
 */

func TestJaccount(t *testing.T) {
	port := rand.Int63n(55535) + 10000

	// start http server
	client := JAccountClient{
		Config: oauth2.Config{
			ClientID:     "FEnnacecAa6yAPWLp3uPjGub",
			ClientSecret: "",
			Scopes:       []string{"lessons", "classes"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://jaccount.sjtu.edu.cn/oauth2/authorize",
				TokenURL: "https://jaccount.sjtu.edu.cn/oauth2/token",
			},
			RedirectURL: fmt.Sprintf("http://localhost:%d/callback", port),
		},
	}

	url, state := client.GetAuthUrl()
	err := openBrowser(url)
	if err != nil {
		t.Error(err)
	}

	http.HandleFunc("/callback", func(writer http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		stateQuery, present := query["state"]
		if !present || len(stateQuery) == 0 || stateQuery[0] != state {
			t.Errorf("error while checking state from oauth callback")
		}

		code, present := query["code"]
		if !present || len(code) == 0 {
			t.Errorf("error while getting code from oauth callback")
		}

		err := client.Auth(code[0])
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("got access_token: %s", client.Token.AccessToken)

		profile, err := client.GetProfile()
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("got profile: %s", profile.Name)

		os.Exit(0)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
