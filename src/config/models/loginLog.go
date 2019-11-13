package models

type LoginLog struct {
	LogId     uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserId    uint64 `gorm:"NOT NULL"`
	DeptId    uint64 `gorm:"NOT NULL"`
	Ip        string `gorm:"NOT NULL"`
	Device    string `gorm:"NOT NULL"`
	LoginType uint8  `gorm:"NOT NULL;DEFAULT:1"`
	CreatedAt int64  `gorm:"NOT NULL"`
}

func (LoginLog) TableName() string {
	return "login_log"
}
