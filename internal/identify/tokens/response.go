package tokens

import (
	"time"
)

// Response is struct
type Response struct {
	Access *Access `json:"access"`
}

// Access is struct
type Access struct {
	Token *Token `json:"token"`
}

// Token is struct
type Token struct {
	ID      string    `json:"id"`
	Expires time.Time `json:"expires"`
}
