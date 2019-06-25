package db


type User struct {
	ID uint `json:"id,omitempty"gorm:"primary_key;AUTO_INCREMENT"`

	Username string `json:"username"gorm:"unique;not null"`

	Password string `json:"password,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`

	Sub string `json:"sub,omitempty"gorm:"-"`

	PreferredUsername string `json:"preferred_username,omitempty"gorm:"-"`
}
