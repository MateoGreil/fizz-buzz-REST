package models

import (
	"mateogreil/fizz-buzz-REST/cmd/config"

	"gorm.io/gorm"
)

type FizzBuzz struct {
	gorm.Model `json:"-";`
	Int1       int    `gorm:"not null;`
	Int2       int    `gorm:"not null;"`
	Limit      int    `gorm:"not null;"`
	Str1       string `gorm:"not null;"`
	Str2       string `gorm:"not null;"`
}

func MaxRequestedFizzBuzz(fizzBuzz *FizzBuzz) {
	config.Db().
		Joins("JOIN queries ON queries.fizz_buzz_id = fizz_buzzs.id").
		Order("count_queries desc").
		Select("fizz_buzzs.*, COUNT(queries.id) AS count_queries").
		Group("fizz_buzzs.id").
		First(&fizzBuzz)
}

func FindOrCreateFizzBuzz(fizzBuzz *FizzBuzz) {
	if config.Db().Where(fizzBuzz).First(fizzBuzz).RowsAffected == 0 {
		config.Db().Create(&fizzBuzz)
	}
}
