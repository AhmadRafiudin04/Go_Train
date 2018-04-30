package businessprocess

import (
	"github.com/gin-gonic/gin"

	"github.com/GinEduTren/dataaccess"
)

func AllJabatan(c *gin.Context) {
	dataaccess.AllGetJabatans(c)
}

func GetJabatan(c *gin.Context) {
	dataaccess.GetJabatan(c)
}

func NewJabatan(c *gin.Context) {
	dataaccess.NewJabatan(c)
}

func DeleteJabatan(c *gin.Context) {
	dataaccess.DeleteJabatan(c)
}
func UpdateJabatan(c *gin.Context) {
	dataaccess.UpdateJabatan(c)
}
