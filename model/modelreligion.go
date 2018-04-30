package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Religion struct {
	gorm.Model
	ReligionName string
}
