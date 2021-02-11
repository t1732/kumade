package tokens

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/t1732/kumade/internal/identify"
)

// API endpoint path
const ENDPOINT_PATH = "/v2.0/tokens"

// GetToken return tokens Response
func GetToken() (*Response, error) {
	credentials := &Credentials{
		Username: os.Getenv("CONOHA_USER_NAME"),
		Password: os.Getenv("CONOHA_PASSWORD"),
	}

	auth := &Auth{
		Credentials: credentials,
		TenantID:    os.Getenv("CONOHA_TENANT_ID"),
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

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
