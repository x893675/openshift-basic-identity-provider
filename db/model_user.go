package db

type User struct {

	Id string `json:"id,omitempty"`

	Username string `json:"username"`

	Password string `json:"password"`

	Email string `json:"email,omitempty"`

	Sub string `json:"sub,omitempty"`

	Name string `json:"name,omitempty"`

	PreferredUsername string `json:"preferred_username,omitempty"`
}