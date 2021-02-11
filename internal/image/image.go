package image

import (
	"fmt"
	"net/http"
	"time"
)

type VPCImage struct {
	Checksum string `json:"checksum"`
	ID string `json:"id"`
	Status string `json:"status"`
	Name string `json:"name"`
	Size int `json:"size"`
	Visibility string `json:"visibility"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const API_HOST string = "https://image-service.tyo1.conoha.io/"
const ENDPOINT_PATH string = "/v2/images"

// curl -i -X DELETE -H 'Content-Type: application/json' -H "Accept: application/json" -H "X-Auth-Token: :token" "https://image-service.tyo1.conoha.io/v2/images/:image_id
func (image *VPCImage) Delete(token string) (error) {
	req, err := http.NewRequest("DELETE", API_HOST + ENDPOINT_PATH + "/" + image.ID, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", token)

  client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 204 {
		return fmt.Errorf("Error: Delete status code: %d", res.StatusCode)
	}

	return nil
}
