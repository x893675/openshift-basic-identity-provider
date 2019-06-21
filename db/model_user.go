package db

import "time"

type User struct {
	ID uint `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`

	Username string `json:"username" gorm:"unique;not null`

	Password string `json:"password,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`

	UpdatedAt time.Time `json:"name,omitempty"`

	CreatedAt time.Time `json:"name,omitempty"`

	DeletedAt *time.Time `json:"name,omitempty"`
}
