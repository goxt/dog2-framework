package models

type Auth struct {
	AuthKey  string `gorm:"PRIMARY_KEY"`
	AuthName string `gorm:"NOT NULL"`
}

func (Auth) TableName() string {
	return "auth"
}
