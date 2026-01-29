package controllers

import (
	"go_gin/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControlerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesService.CreateNotesService(),
		})
	}
}
