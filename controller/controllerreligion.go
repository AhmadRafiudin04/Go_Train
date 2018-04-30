package controller

import (
	"github.com/GinEduTren/businessprocess"
	"github.com/gin-gonic/gin"
)

// Main Controll untuk routing Link API
func ReligionControll() {
	r := gin.Default()

	r.GET("/api/religion", businessprocess.AllReligion)
	r.GET("/api/religion/:id", businessprocess.GetReligion)
	r.POST("api/religion", businessprocess.NewReligion)
	r.PUT("/api/religion/:id", businessprocess.UpdateReligion)
	r.DELETE("/api/religion/:id", businessprocess.DeleteReligion)
	r.Run(":8080")

}
