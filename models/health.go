package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"health/global"
	"time"
)

type Base struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Health struct {
	Base
	Uid    uint32  `gorm:"index" form:"uid" json:"uid"`
	Weight float32 `form:"weight" json:"weight" binding:"required"`
	Waist  float32 `form:"waist" json:"waist"`
	Bust   float32 `form:"bust" json:"bust"`
	Hip    float32 `form:"hip" json:"hip"`
	Thigh  float32 `form:"thigh" json:"thigh"`
	Arm    float32 `form:"arm" json:"arm"`
	Day    string  `gorm:"type:date;index" form:"day" json:"day" binding:"required"`
}

type Cond struct {
	Id    uint   `form:"id"`
	Uid   uint32 `form:"uid"`
	Start string `form:"start_time"`
	End   string `form:"end_time"`
}

func Create(data Health) (*Health, error) {
	var h Health
	if result := global.Db.Where(&Health{Uid: data.Uid, Day: data.Day}).First(&h); result.RowsAffected > 0 {
		return nil, errors.New("当天记录已存在")
	}
	if result := global.Db.Create(&data); result.RowsAffected == 0 {
		return nil, errors.New("记录添加失败")
	}
	return &data, nil
}

func Del(data Cond) error {
	var h Health
	if result := global.Db.First(&h, data.Id); result.RowsAffected == 0 {
		return errors.New("记录不存在")
	}
	if result := global.Db.Where(&Health{Uid: data.Uid}).Delete(&Health{}, data.Id); result.RowsAffected == 0 {
		return errors.New("记录删除失败")
	}
	return nil
}

func List(data Cond) ([]Health, error) {
	var ret []Health
	db := global.Db.Model(&Health{})
	fmt.Printf("uid :%d", data.Uid)
	db = db.Where(&Health{Uid: data.Uid})
	if data.Start != "" && data.End != "" {
		db = db.Where("day between ? and ?", data.Start, data.End)
	}
	db.Find(&ret)
	return ret, nil
}
