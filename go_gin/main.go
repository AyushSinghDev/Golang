package main

import (
	"go_gin/controllers"

	"github.com/gin-gonic/gin"
)

// type CreateUserRequest struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type UpdateUserRequest struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required,min=6"`
// }

// type PatchUserRequest struct {
// 	Email    *string `json:"email"`
// 	Password *string `json:"password"`
// }

func main() {
	router := gin.Default()

	//Test pings
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong from go-gin",
	// 	})
	// })

	// //Get
	// r.GET("/me/:id/:newId", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	newID := c.Param("newId")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"user_id":    id,
	// 		"user_newid": newID,
	// 	})
	// })

	// //Create
	// r.POST("/me", func(c *gin.Context) {
	// 	var req CreateUserRequest
	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"message": "user created successfully",
	// 		"email":   req.Email,
	// 	})
	// })

	// //Update
	// r.PUT("/me/:id", func(c *gin.Context) {
	// 	var req UpdateUserRequest
	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "user replaced successfully",
	// 		"email":   req.Email,
	// 	})
	// })

	// //Patch
	// r.PATCH("/me/:id", func(c *gin.Context) {
	// 	var req PatchUserRequest
	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	//At least one field required
	// 	if req.Email == nil && req.Password == nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": "at least one field must be provided",
	// 		})
	// 		return
	// 	}
	// 	if req.Email != nil {
	// 		// update email
	// 	}
	// 	if req.Password != nil {
	// 		// update password
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "user updated successfully",
	// 	})
	// })

	// // DELETE
	// r.DELETE("/me/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"id":      id,
	// 		"message": "user deleted successfully",
	// 	})
	// })

	// Server start
	// if err := r.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	notesController := &controllers.NotesController{}
	notesController.InitNotesControlerRoutes(router)

	router.Run()
}
