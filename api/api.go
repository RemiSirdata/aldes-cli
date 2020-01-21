package api

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Endpoint = "https://aldesiotsuite-aldeswebapi.azurewebsites.net"
)

var (
	EndpointOauth2 = oauth2.Endpoint{
		AuthURL:  Endpoint + "/oauth2/authorize",
		TokenURL: Endpoint + "/oauth2/token",
	}
	Timeout = time.Minute
)

func New(login string, password string) (*Aldes, error) {
	a := Aldes{}
	err := a.auth(login, password)
	return &a, err
}

type Aldes struct {
	Token *oauth2.Token
}

func (a *Aldes) auth(login string, password string) error {
	ctx := context.Background()
	conf := &oauth2.Config{
		Scopes:   []string{"offline_access"},
		Endpoint: EndpointOauth2,
	}
	token, err := conf.PasswordCredentialsToken(ctx, login, password)
	if err != nil {
		return err
	}
	a.Token = token
	return nil
}

func (a *Aldes) GetProducts() ([]Product, error) {
	var products []Product
	err := a.Get(&products, "/aldesoc/v2/users/me/products")
	return products, err
}

func (a *Aldes) Get(model interface{}, path string, args ...interface{}) error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(Endpoint+path, args...), nil)
	if err != nil {
		return err
	}
	a.Token.SetAuthHeader(req)
	client := http.Client{Timeout: Timeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("invalid status code '%d' response %s", resp.StatusCode, body)
	}
	return json.NewDecoder(resp.Body).Decode(&model)

}
