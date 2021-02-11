package images

import "github.com/t1732/kumade/internal/image"

type Response struct {
	Images *[]image.VPCImage `json:"images"`
}
