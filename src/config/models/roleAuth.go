package models

type RoleAuth struct {
	RoleKey string `gorm:"PRIMARY_KEY"`
	AuthKey string `gorm:"PRIMARY_KEY"`
}

func (RoleAuth) TableName() string {
	return "role_auth"
}
