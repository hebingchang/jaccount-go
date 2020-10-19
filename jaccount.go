package jaccount

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/imroc/req"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/oauth2"
)

/**
* @Author: He Bingchang
* @Date: 2020/9/7
* @Description:
 */

type JAccountClient struct {
	oauth2.Config
	Token *oauth2.Token
}

func (client *JAccountClient) GetAuthUrl() (string, string) {
	state := uuid.NewV4().String()
	url := client.AuthCodeURL(state)
	return url, state
}

func (client *JAccountClient) Auth(code string) error {
	ctx := context.Background()

	token, err := client.Exchange(ctx, code)
	if err != nil {
		return err
	}

	client.Token = token
	return nil
}

func (client *JAccountClient) GetProfile() (*Profile, error) {
	if client.Token == nil {
		return nil, errors.New("haven't exchange access token")
	}

	param := req.Param{
		"access_token": client.Token.AccessToken,
	}

	r, err := req.Get("https://api.sjtu.edu.cn/v1/me/profile", param)
	if err != nil {
		return nil, err
	}

	var resp profileResp
	_ = json.Unmarshal(r.Bytes(), &resp)
	return &resp.Entities[0], nil
}

func (client *JAccountClient) GetCourse(code string) (*Lesson, error) {
	if client.Token == nil {
		return nil, errors.New("haven't exchange access token")
	}

	param := req.Param{
		"access_token": client.Token.AccessToken,
	}

	r, err := req.Get("https://api.sjtu.edu.cn/v1/lesson/"+code, param)
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
	if client.Token == nil {
		return nil, errors.New("haven't exchange access token")
	}

	param := req.Param{
		"access_token": client.Token.AccessToken,
	}

	url := "https://api.sjtu.edu.cn/v1/me/lessons"
	if len(semester) > 0 {
		url += "/" + semester[0]
	}

	r, err := req.Get(url, param)
	if err != nil {
		return nil, err
	}

	var resp lessonsResp
	_ = json.Unmarshal(r.Bytes(), &resp)
	return resp.Entities, nil
}
