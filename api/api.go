package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Api(db *gorm.DB) {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		req := Request{}

		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			log.Fatal(err)
		}

		for _, lang := range req.getAllLanguages() {
			if err := db.Create(&lang); err.Error != nil {
				log.Println(err)
			}
		}

		c.JSON(http.StatusOK, req)
	})
	r.Run()
}
