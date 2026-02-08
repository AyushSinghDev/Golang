package services

import (
	"fmt"
	internal "go_gin/internal/database/model"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{Id: 1, Name: "Note 1"},
		{Id: 2, Name: "Note 2"},
	}
	return data
}

func (n *NotesService) CreateNotesService() Note {
	data := Note{
		Id: 3, Name: "Note 3",
	}

	result := n.db.Create(&internal.Notes{
		Id:     1,
		Title:  "Notes",
		Status: true,
	})
	if result.Error != nil {
		fmt.Print(result.Error)
	}

	return data
}
