package db

type User struct {
	Id string `json:"id,omitempty"`

	Username string `json:"username"`

	Password string `json:"password,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`
}
