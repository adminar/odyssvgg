package model

type UserTable struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	LastLogin string `json:"last_login"`
	AuthToken string `json:"auth_token"`
	Type      string `json:"type"`
}

func (UserTable) TableName() string {
	return "users"
}
