package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model  //Id gen update del
	CateId      string
	User        string `json:"user" gorm:"type:varchar(255);not null"`
	Name        string `json:"name" gorm:"type:varchar(50);not null"`
	Picture     string `json:"picture" gorm:"type:varchar(1024);not null"`
	Price       string `json:"price" gorm:"type:float;not null"`
	Description string `json:"desc" gorm:"type:varchar(255);not null"`
	IsSold      bool   `json:"is_sold" gorm:"type:boolean;not null"`
}
