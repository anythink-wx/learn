package dao

import (
	"gorm.io/gorm"
	"learn/week2/g"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// GetInfo dao层，获取最
func GetInfo(id int) (*Product, error) {
	var pd Product
	err := g.Db.First(&pd, id).Error
	return &pd, err
}
