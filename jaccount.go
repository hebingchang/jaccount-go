package jaccount

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
	"net/http"
)

/**
* @Author: He Bingchang
* @Date: 2020/9/7
* @Description:
 */

type JAccountAuth struct {
	oauth2.Config
}

func (client *JAccountAuth) GetAuthUrl() (string, string) {
	state := uuid.NewV4().String()
	url := client.AuthCodeURL(state)
	return url, state
}

func (client *JAccountAuth) Auth(code string) (*oauth2.Token, error) {
	ctx := context.Background()
	return client.Exchange(ctx, code)
}

func (client *JAccountAuth) AuthWithRetUrl(code string, retUrl string) (*oauth2.Token, error) {
	ctx := context.Background()
	config := client.Config
	config.RedirectURL = retUrl
	return config.Exchange(ctx, code)
}

func (client *JAccountAuth) GetClient(token *oauth2.Token) *http.Client {
	ctx := context.Background()
	return client.Client(ctx, token)
}

type JAccountClient struct {
	Client *http.Client
}

func (client *JAccountClient) GetProfile() (*Profile, error) {
	req.SetClient(client.Client)

	r, err := req.Get("https://api.sjtu.edu.cn/v1/me/profile")
	if err != nil {
		return nil, err
	}

	var resp profileResp
	_ = json.Unmarshal(r.Bytes(), &resp)
	return &resp.Entities[0], nil
}

func (client *JAccountClient) GetCourse(code string) (*Lesson, error) {
	req.SetClient(client.Client)

	r, err := req.Get("https://api.sjtu.edu.cn/v1/lesson/" + code)
	if err != nil {
		return nil, err
	}

	var resp lessonsResp
	_ = json.Unmarshal(r.Bytes(), &resp)

	if resp.Errno == 0 {
		return &resp.Entities[0], nil
	} else {
		return nil, errors.New(resp.Error)
	}
}

func (client *JAccountClient) GetLessons(semester ...string) ([]Lesson, error) {
	req.SetClient(client.Client)

	url := "https://api.sjtu.edu.cn/v1/me/lessons"
	if len(semester) > 0 {
		url += "/" + semester[0]
	}

	r, err := req.Get(url)
	if err != nil {
		return nil, err
	}

	var resp lessonsResp
	_ = json.Unmarshal(r.Bytes(), &resp)
	return resp.Entities, nil
}
