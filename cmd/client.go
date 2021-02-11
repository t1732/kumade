package cmd

import (
	"fmt"
	"log"

	"github.com/t1732/kumade/internal/identify/tokens"
	"github.com/t1732/kumade/internal/image"
	"github.com/t1732/kumade/internal/image/images"
)

func GetToken() *tokens.Token {
	identity, err := tokens.GetToken()
	if err != nil {
		fmt.Println("[!] " + err.Error())
		log.Fatal(err)
	}

	return identity.Access.Token
}

func GetVPCImages(token string) *[]image.VPCImage {
	vpcImages, err := images.GetVPCImages(token)
	if err != nil {
		fmt.Println("[!] " + err.Error())
		log.Fatal(err)
	}

	return vpcImages
}
