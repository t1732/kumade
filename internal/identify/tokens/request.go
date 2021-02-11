package tokens

// RequestParams is struct
type RequestParams struct {
	Auth *Auth `json:"auth"`
}

// Auth is Struct
type Auth struct {
	Credentials *Credentials `json:"passwordCredentials"`
	TenantID    string       `json:"tenantId"`
}

// Credentials is struct
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
