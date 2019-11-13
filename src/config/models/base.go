package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	CreatedId uint64 `gorm:"NOT NULL"`
	CreatedAt int64  `gorm:"NOT NULL"`
	DeletedAt *int64 `gorm:"DEFAULT:NULL"`
	BaseOrm   `gorm:"-"`
}

type BaseOrm struct {
	theOperatorId  uint64 `gorm:"-"`
	theOperationAt int64  `gorm:"-"`
	theOperatorDel bool   `gorm:"-"`
}

const (
	SystemUserId   = 0
	SystemUserName = "系统"
	SystemDeptId   = 0
	SystemDeptName = "系统"
)

func (base *BaseOrm) SetCreator(createdId uint64, createdAt ...int64) {
	if len(createdAt) > 0 {
		base.theOperationAt = createdAt[0]
	} else {
		base.theOperationAt = time.Now().Unix()
	}

	base.theOperatorId = createdId
}

func (base *BaseOrm) BeforeCreate(scope *gorm.Scope) error {
	if base.theOperatorId == 0 {
		base.SetCreator(SystemUserId)
	}
	_ = scope.SetColumn("CreatedId", base.theOperatorId)
	_ = scope.SetColumn("CreatedAt", base.theOperationAt)
	return nil
}

func (base *BaseOrm) BeforeUpdate(scope *gorm.Scope) error {
	return nil
}

func (base *BaseOrm) BeforeDelete(scope *gorm.Scope) error {
	if base.theOperatorId == 0 {
		base.SetCreator(SystemUserId)
	}
	base.theOperatorDel = true
	var sql = "UPDATE `" + scope.TableName() + "` SET " +
		"deleted_at = ? WHERE `" + scope.PrimaryKey() + "` = ?"
	rst := scope.DB().Exec(sql, base.theOperationAt, scope.PrimaryKeyValue())
	if rst.Error != nil {
		panic(rst.Error)
	}
	return nil
}
