package dataaccess

import (
	"fmt"
	"net/http"

	"github.com/GinEduTren/config"
	"github.com/GinEduTren/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//var Db *gorm.DB
//var Err error
var jabatans []model.Jabatan

// Get Data Religion for All Data
func AllGetJabatans(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Jabatan{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	err := db.Find(&jabatans).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the data"})
		fmt.Println(err)
	} else {
		//c.JSON(200, religions)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &jabatans})
	}
}

// get data religion bby Filter ID
func GetJabatan(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Jabatan{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var jabatan model.Jabatan
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&jabatan).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the data"})
		fmt.Println(err)
	} else {
		//c.JSON(200, religion)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &jabatan})
	}

}

// Insert Data Religion
func NewJabatan(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Jabatan{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var jabatan model.Jabatan

	c.BindJSON(&jabatan)
	err := db.Create(&jabatan)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to insert the data"})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &jabatan})
	}

}

// Delete Data Religion by ID
func DeleteJabatan(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Jabatan{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var jabatan model.Jabatan
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).Delete(&jabatan).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the data"})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &jabatan})
	}

}

// Update Data Religion By ID and Set Field based on JSon Data sent
func UpdateJabatan(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Jabatan{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var jabatan model.Jabatan
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&jabatan).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.BindJSON(&jabatan)
		db.Save(&jabatan)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &jabatan})
	}
}
