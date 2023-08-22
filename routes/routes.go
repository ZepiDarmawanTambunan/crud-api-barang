package routes

import (
	"crud-api-barang/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	barangGroup := r.Group("/barang")
	{
		barangGroup.GET("/", func(c *gin.Context) { controllers.GetAllBarang(c, db) })
		barangGroup.GET("/:id", func(c *gin.Context) { controllers.GetBarangByID(c, db) })
		barangGroup.POST("/", func(c *gin.Context) { controllers.CreateBarang(c, db) })
		barangGroup.PUT("/:id", func(c *gin.Context) { controllers.UpdateBarang(c, db) })
		barangGroup.DELETE("/:id", func(c *gin.Context) { controllers.DeleteBarang(c, db) })
	}

	return r
}
