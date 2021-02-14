package conoha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

const (
	imageAPIHost            = "https://image-service.tyo1.conoha.io/"
	imageImagesEndpointPath = "/v2.0/images"
)

type imageAPIData struct {
	token    string
	tenantId string
}

type imagesResponse struct {
	Images *[]VMImage `json:"images"`
}

type VMImage struct {
	Checksum   string    `json:"checksum"`
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	Name       string    `json:"name"`
	Size       int       `json:"size"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type imagesOption func(*imagesSearchOption)

type imagesSearchOption struct {
	Status string
}

func Image(token string) *imageAPIData {
	return &imageAPIData{
		tenantId: viper.GetString("tenant_id"),
		token:    token,
	}
}

func ImageStatus(status string) imagesOption {
	return func(option *imagesSearchOption) {
		option.Status = status
	}
}

//	VM イメージ一覧取得
//
//	ステータスを指定して取得する場合
//	GetImages(ImageStatus("active"))
func (data *imageAPIData) GetImages(options ...imagesOption) (*[]VMImage, error) {
	u, err := url.Parse(imageAPIHost + imageImagesEndpointPath)
	if err != nil {
		return nil, err
	}

	searchOption := &imagesSearchOption{Status: "active"}

	q := u.Query()
	q.Set("owner", data.tenantId)
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
		return nil, fmt.Errorf("Image API images request error: %s", body)
	}

	response := &imagesResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response.Images, nil
}

// VM イメージの削除
// curl -i -X DELETE -H 'Content-Type: application/json' -H "Accept: application/json" -H "X-Auth-Token: :token" "https://image-service.tyo1.conoha.io/v2/images/:image_id
func (data *imageAPIData) Delete(imageID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintln(imageAPIHost, imageImagesEndpointPath, "/", imageID), nil)
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
