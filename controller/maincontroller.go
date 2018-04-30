package controller

import (
	"github.com/GinEduTren/businessprocess"
	"github.com/gin-gonic/gin"
)

var err error

// untuk Controller Utama, yang membagikan jenis2 controller yang akan digunakan
func MainController() {
	r := gin.Default()
	//ReligionControll()
	r.GET("/api/religion", businessprocess.AllReligion)
	r.GET("/api/religion/:id", businessprocess.GetReligion)
	r.POST("api/religion", businessprocess.NewReligion)
	r.PUT("/api/religion/:id", businessprocess.UpdateReligion)
	r.DELETE("/api/religion/:id", businessprocess.DeleteReligion)

	//JabatanControll()
	r.GET("/api/jabatan", businessprocess.AllJabatan)
	r.GET("/api/jabatan/:id", businessprocess.GetJabatan)
	r.POST("api/jabatan", businessprocess.NewJabatan)
	r.PUT("/api/jabatan/:id", businessprocess.UpdateJabatan)
	r.DELETE("/api/jabatan/:id", businessprocess.DeleteJabatan)

	// running route
	r.Run(":8080")

}
