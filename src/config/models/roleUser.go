package models

type RoleUser struct {
	RoleKey string `gorm:"PRIMARY_KEY"`
	UserId  string `gorm:"PRIMARY_KEY"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
