package controller

import (
	"github.com/GinEduTren/businessprocess"
	"github.com/gin-gonic/gin"
)

// Main Controll untuk routing Link API
func JabatanControll() {
	r := gin.Default()

	r.GET("/api/jabatan", businessprocess.AllJabatan)
	r.GET("/api/jabatan/:id", businessprocess.GetJabatan)
	r.POST("api/jabatan", businessprocess.NewJabatan)
	r.PUT("/api/jabatan/:id", businessprocess.UpdateJabatan)
	r.DELETE("/api/jabatan/:id", businessprocess.DeleteJabatan)
	//r.Run(":8080")

}
