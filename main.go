package main

import (
	"fmt"

	"github.com/t1732/kumade/cmd"
)

func main() {
	token := cmd.GetToken()
	images := cmd.GetVPCImages(token.ID)

	fmt.Println(images)
}
