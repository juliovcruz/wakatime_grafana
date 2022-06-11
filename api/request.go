package api

import (
	"github.com/google/uuid"
	"main.go/model"
	"time"
)

type Request struct {
	User RequestUser  `json:"user" binding:"required"`
	Days []RequestDay `json:"days" binding:"required"`
}

type RequestUser struct {
	Email string `json:"email" binding:"required"`
}

type RequestDay struct {
	Editors   []RequestStruct `json:"editors" binding:"required"`
	Languages []RequestStruct `json:"languages" binding:"required"`
	Machines  []RequestStruct `json:"machines" binding:"required"`
	Date      RequestDate     `json:"date" binding:"required" time_format:"2006-01-02"`
}

type RequestStruct struct {
	Hours   int    `json:"hours" binding:"required"`
	Minutes int    `json:"minutes" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Text    string `json:"text" binding:"required"`
}

func (r Request) getAllLanguages() []model.Language {
	var languages []model.Language

	for _, day := range r.Days {
		for _, lang := range day.Languages {
			minutes := lang.Minutes + (lang.Hours * 60)

			if minutes != 0 {
				languages = append(languages, model.Language{
					ID:   uuid.Must(uuid.NewRandom()),
					Date: time.Time(day.Date),
					RequestStruct: model.RequestStruct{
						Hours: float32(minutes) / 60,
						Name:  lang.Name,
						Email: r.User.Email,
					},
				})
			}
		}
	}

	return languages
}
