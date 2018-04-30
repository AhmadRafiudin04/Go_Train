package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Jabatan struct {
	gorm.Model
	JabatanName string `json:"JabatanName"`
}
