package db

type User struct {
	ID uint `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`

	Username string `json:"username" gorm:"unique;not null"`

	Password string `json:"password,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`

	Sub string `json:"sub,omitempty" gorm:"-"`

	PreferredUsername string `json:"preferred_username,omitempty" gorm:"-"`

	Role uint `json:"role,omitempty" gorm:"column:role"` //1 is admin , 2 is common user
}

type UpdatePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserEmailParam struct {
	Email string `json:"email"`
}

type MailResetPassword struct {
	Password string `json:"password"`
}
