package conoha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	networkAPIHost                    = "https://networking.tyo1.conoha.io"
	networkSecurityGroupsEndpointPath = "/v2.0/security-groups"
)

type securityGroupsResponse struct {
	SecurityGroups *[]SecurityGroup `json:"security_groups"`
}

type SecurityGroup struct {
	ID                 string               `json:"id"`
	Name               string               `json:"name"`
	Description        string               `json:"description"`
	SecurityGroupRules *[]SecurityGroupRule `json:"security_group_rules"`
}

type SecurityGroupRule struct {
	ID              string `json:"id"`
	Direction       string `json:"direction"`
	Ethertype       string `json:"ethertype"`
	PortRangeMax    int    `json:"port_range_max"`
	PortRangeMin    int    `json:"port_range_min"`
	SecurityGroupId string `json:"security_group_id"`
}

type networkAPIData struct {
	token string
	url   *url.URL
}

func Network(token string) *networkAPIData {
	u, err := url.Parse(networkAPIHost)
	if err != nil {
		log.Fatal(err)
	}

	return &networkAPIData{
		token: token,
		url:   u,
	}
}

//	セキュリティグループ取得
func (data *networkAPIData) GetSecurityGroups() (*[]SecurityGroup, error) {
	data.url.Path = networkSecurityGroupsEndpointPath

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
		return nil, fmt.Errorf("Network API security groups request error: %s", body)
	}

	response := &securityGroupsResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response.SecurityGroups, nil
}

// セキュリティグループの削除
func (data *networkAPIData) DeleteSecurityGroup(securityGroupID string) error {
	data.url.Path = path.Join(networkSecurityGroupsEndpointPath, securityGroupID)

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
		return fmt.Errorf("Netrowk API security group delete request error: %d", res.StatusCode)
	}

	return nil
}
