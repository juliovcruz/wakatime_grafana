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

		languages := req.getAllLanguages()

		for _, lang := range languages {
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

		c.JSON(http.StatusOK, languages)
	})

	r.GET("/backup/sql", func(c *gin.Context) {
		var languages []model.Language
		result := db.Find(&languages)
		if result.Error != nil {
			log.Fatal(result.Error)
		}

		c.String(http.StatusOK, model.ConvertToBackupSQL(languages))
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
