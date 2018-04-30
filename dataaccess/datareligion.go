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

var db *gorm.DB
var err error
var religions []model.Religion

// Religion Model struct for gorm
//type Religion struct {
//	gorm.Model
//	ReligionName string
//}

// Get Data Religion for All Data
func AllGetReligions(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Religion{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	err := db.Find(&religions).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the data"})
		fmt.Println(err)
	} else {
		//c.JSON(200, religions)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &religions})
	}
}

// get data religion bby Filter ID
func GetReligion(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Religion{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var religion model.Religion
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&religion).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to read the data"})
		fmt.Println(err)
	} else {
		//c.JSON(200, religion)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &religion})
	}

}

// Insert Data Religion
func NewReligion(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Religion{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var religion model.Religion

	c.BindJSON(&religion)
	err := db.Create(&religion)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to insert the data"})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &religion})
	}

}

// Delete Data Religion by ID
func DeleteReligion(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Religion{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var religion model.Religion
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).Delete(&religion).Error
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": "Failed to delete the data"})
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &religion})
	}

}

// Update Data Religion By ID and Set Field based on JSon Data sent
func UpdateReligion(c *gin.Context) {
	psqlInfo := config.ConnString()
	db, err = gorm.Open("postgres", psqlInfo)
	db.AutoMigrate(&model.Religion{})
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var religion model.Religion
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&religion).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.BindJSON(&religion)
		db.Save(&religion)

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &religion})
	}
}
