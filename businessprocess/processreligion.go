package businessprocess

import (
	"github.com/gin-gonic/gin"

	"github.com/GinEduTren/dataaccess"
)

func AllReligion(c *gin.Context) {
	dataaccess.AllGetReligions(c)
}

func GetReligion(c *gin.Context) {
	dataaccess.GetReligion(c)
}

func NewReligion(c *gin.Context) {
	dataaccess.NewReligion(c)
}

func DeleteReligion(c *gin.Context) {
	dataaccess.DeleteReligion(c)
}
func UpdateReligion(c *gin.Context) {
	dataaccess.UpdateReligion(c)
}
