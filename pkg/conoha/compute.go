package conoha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/structs"
	"github.com/spf13/viper"
)

const (
	computeAPIHost               = "https://compute.tyo1.conoha.io/"
	computeServersEndpointFormat = "/v2/%s/servers"
	computeDeleteEndpointFormat  = "/v2/%s/servers/%s"
)

type computeAPIData struct {
	token    string
	tenantId string
	url      *url.URL
}

type serversResponse struct {
	Servers *[]Server `json:"servers"`
}

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type serversOption func(*serversSearchOption)

type serversSearchOption struct {
	Name   string
	Status string
}

func ServerName(name string) serversOption {
	return func(option *serversSearchOption) {
		option.Name = name
	}
}

func ServerStatus(status string) serversOption {
	return func(option *serversSearchOption) {
		option.Status = status
	}
}

func Compute(token string) *computeAPIData {
	u, err := url.Parse(computeAPIHost)
	if err != nil {
		log.Fatal(err)
	}

	return &computeAPIData{
		token:    token,
		tenantId: viper.GetString("tenant_id"),
		url:      u,
	}
}

func (data *computeAPIData) GetServers(options ...serversOption) (*[]Server, error) {
	data.url.Path = fmt.Sprintf(computeServersEndpointFormat, data.tenantId)

	searchOption := &serversSearchOption{}
	for _, option := range options {
		option(searchOption)
	}

	q := data.url.Query()
	q.Set("owner", data.tenantId)
	opts := structs.Map(searchOption)
	for k, v := range opts {
		if v.(string) != "" {
			q.Set(k, v.(string))
		}
	}
	data.url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", data.url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", data.token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Compute API servers request error: %s", body)
	}

	response := &serversResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response.Servers, nil
}

//	VM 削除
//	DELETE /v2/​{tenant_id}​/servers/​{server_id}​
func (data *computeAPIData) DeleteServer(serverID string) error {
	data.url.Path = fmt.Sprintf(computeDeleteEndpointFormat, data.tenantId, serverID)
	req, err := http.NewRequest("DELETE", data.url.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", data.token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return fmt.Errorf("Image API delete request error: %d", res.StatusCode)
	}

	return nil
}
