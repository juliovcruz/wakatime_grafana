package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"main.go/model"
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
				var dbResult model.Language
				if err := db.Where("date = ? AND name = ?", lang.Date, lang.Name).Find(&dbResult); err.Error != nil {
					log.Fatal(err)
				}

				if lang.Hours > dbResult.Hours {
					db.Model(&dbResult).Update("hours", lang.Hours)
				}
			}
		}

		c.JSON(http.StatusOK, req)
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
