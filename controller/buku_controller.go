package controller

import (
	"biodata/database"
	"biodata/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBuku(c *gin.Context) {

	var buku []model.Buku
	var err error

	buku, err = model.ReadBuku(database.DB)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": buku,
	})

}

func CreateBuku(c *gin.Context) {
	var buku model.Buku

	if err := c.ShouldBind(&buku); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	err := model.CreateBuku(database.DB, buku)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Buku created successfully",
		"data":    buku,
	})

}

func UpdateBuku(c *gin.Context) {
	var GetData model.Buku

	if err := c.ShouldBind(&GetData); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	var id = c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	buku := model.Buku{
		ID:          idConv,
		Judul:       GetData.Judul,
		Penulis:     GetData.Penulis,
		TahunTerbit: GetData.TahunTerbit,
	}

	hasil, err := model.UpdateBuku(database.DB, idConv, buku)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Buku updated successfully",
		"data":    hasil,
	})

}

func DeleteBuku(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	if id == "" {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	err = model.DeleteBuku(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Buku deleted successfully",
	})
}

func GetBukuById(c *gin.Context) {
	id := c.Param("id")
	idConv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	buku, err := model.GetBukuById(database.DB, idConv)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request, invalid data",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": buku,
	})
}
