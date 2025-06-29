package route

import (
	"biodata/controller"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.GET("/petugas", controller.GetAllPetugas)
	router.GET("/petugas/:id", controller.GetPetugasById)
	router.POST("/petugas", controller.CreatePetugas)
	router.PUT("/petugas/:id", controller.UpdatePetugas)
	router.DELETE("/petugas/:id", controller.DeletePetugas)

	router.GET("/buku", controller.GetAllBuku)
	router.GET("/buku/:id", controller.GetBukuById)
	router.POST("/buku", controller.CreateBuku)
	router.PUT("/buku/:id", controller.UpdateBuku)
	router.DELETE("/buku/:id", controller.DeleteBuku)

	return router
}
