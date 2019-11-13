package models

type User struct {
	UserId        uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	DeptId        uint64 `gorm:"DEFAULT:0"`
	UserName      string `gorm:"NOT NULL"`
	Account       string `gorm:"NOT NULL"`
	AccountStatus uint8  `gorm:"NOT NULL;DEFAULT:1"`
	Pwd           string `gorm:"NOT NULL"`
	DynamicPwd    string `gorm:"DEFAULT:NULL"`
	DynamicPwdAt  int64  `gorm:"DEFAULT:NULL"`
	Base
}

func (User) TableName() string {
	return "user"
}
