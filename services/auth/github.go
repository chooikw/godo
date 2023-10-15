package authservice

import (
	"errors"
	"fmt"
	"os"

	"encoding/json"
	"io"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GithubUser struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var githubOauthConfig = oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Endpoint:     github.Endpoint,
}

func GithubVerifierUrl() (url string) {
	fmt.Println(os.Getenv("GITHUB_CLIENT_ID"))
	url = githubOauthConfig.AuthCodeURL("state")
	return
}

func GithubVerifyCode(code string) (User, error) {
	var user User
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return user, errors.New("Github verification failed")
	}
	client := githubOauthConfig.Client(oauth2.NoContext, token)
	res, err := client.Get("https://api.github.com/user")
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var githubUser GithubUser
	json.Unmarshal(body, &githubUser)
	user.Id = fmt.Sprintf("github:%v", githubUser.Id)
	user.Name = githubUser.Name
	return user, nil
}
