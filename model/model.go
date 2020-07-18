package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name   string
	Brands []Brand
	Unit   string
}

type Brand struct {
	gorm.Model
	Name      string
	ProductID uint
}

type Purchase struct {
	gorm.Model
	ProductID uint
	BrandID   uint
	Unit      string
	Price     float64
	Quantity  float64
}
