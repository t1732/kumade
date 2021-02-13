package images

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
	"github.com/t1732/kumade/internal/image"
)

const ENDPOINT_PATH string = "/v2/images"

type Option func(*SearchOption)

type SearchOption struct {
	Status string
}

func ImageStatus(status string) Option {
	return func(option *SearchOption) {
		option.Status = status
	}
}

// GetVPCImages return VPCImage list
func GetVPCImages(token string, options ...Option) (*[]image.VPCImage, error) {
	u, err := url.Parse(image.API_HOST + ENDPOINT_PATH)
	if err != nil {
		return nil, err
	}

	searchOption := &SearchOption{Status: "active"}

	q := u.Query()
	q.Set("owner", viper.GetString("tenant_id"))
	for _, option := range options {
		option(searchOption)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", token)
	client := &http.Client{}
	res, err := client.Do(req)
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

	return response.Images, nil
}
