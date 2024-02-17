package model

import "gorm.io/gorm"

type Image struct {
	AltText   string
	ImageLink string
	Title     string
}

type Products struct {
	*gorm.Model
	ID                 int
	Name               string
	Description        string
	Price              float64
	Quantity           int
	MetaTitle          string
	MetaDescription    string
	Number_of_variants int
	Images             []string
}

type ProductVariant struct {
	*gorm.Model
	ID         int
	ProductID  int
	Type       string
	Variantion string
}
