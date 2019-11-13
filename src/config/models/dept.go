package models

type Dept struct {
	DeptId    uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	DeptPid   uint64 `gorm:"NOT NULL"`
	DeptType  uint8  `gorm:"NOT NULL"`
	DeptPtype uint8  `gorm:"NOT NULL"`
	DeptName  string `gorm:"NOT NULL"`
	Sort      uint32 `gorm:"DEFAULT:100"`
	Base
}

type DeptExt struct{}

func (Dept) TableName() string {
	return "dept"
}
