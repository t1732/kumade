package tokens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	"github.com/t1732/kumade/internal/identify"
)

// API endpoint path
const ENDPOINT_PATH = "/v2.0/tokens"

// GetToken return tokens Response
func GetToken() (*Response, error) {
	credentials := &Credentials{
		Username: viper.GetString("user"),
		Password: viper.GetString("password"),
	}

	auth := &Auth{
		Credentials: credentials,
		TenantID:    viper.GetString("tenant_id"),
	}

	requestParams := &RequestParams{
		Auth: auth,
	}

	jsonBytes, err := json.Marshal(requestParams)
	if err != nil {
		return nil, err
	}

	endpoint := identify.API_HOST + ENDPOINT_PATH
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

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
