package models

type Role struct {
	RoleKey  string `gorm:"PRIMARY_KEY"`
	RoleName string `gorm:"NOT NULL"`
	IsSystem bool   `gorm:"NOT NULL;DEFAULT:0"`
}

func (Role) TableName() string {
	return "role"
}
