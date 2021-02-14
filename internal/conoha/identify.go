package conoha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

const (
	identifyAPIHost            = "https://identity.tyo1.conoha.io"
	identifyTokensEndpointPath = "/v2.0/tokens"
)

type identifyAPIData struct {
	user     string
	password string
	tenantId string
}

type tokensRequestParams struct {
	Auth *tokensRequestAuth `json:"auth"`
}

type tokensRequestAuth struct {
	Credentials *tokensRequestCredentials `json:"passwordCredentials"`
	TenantID    string                    `json:"tenantId"`
}

type tokensRequestCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokensResponse struct {
	Access *TokensResponseAccess `json:"access"`
}

type TokensResponseAccess struct {
	Token *Token `json:"token"`
}

type Token struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
}

func Identify() *identifyAPIData {
	return &identifyAPIData{
		user:     viper.GetString("user"),
		password: viper.GetString("password"),
		tenantId: viper.GetString("tenant_id"),
	}
}

//	APIトークン発行
func (data *identifyAPIData) CreateToken() (*tokensResponse, error) {
	credentials := &tokensRequestCredentials{
		Username: data.user,
		Password: data.password,
	}

	auth := &tokensRequestAuth{
		Credentials: credentials,
		TenantID:    data.tenantId,
	}

	requestParams := &tokensRequestParams{
		Auth: auth,
	}

	jsonBytes, err := json.Marshal(requestParams)
	if err != nil {
		return nil, err
	}

	endpoint := identifyAPIHost + identifyTokensEndpointPath
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Token API request error: %s", body)
	}

	response := &tokensResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
