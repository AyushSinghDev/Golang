package controllers

import (
	"gorm/initializers"
	"gorm/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id of URL
	id := c.Param("id")

	//Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with items
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id of URL
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context){
	//Get the id off the URL
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//Respond with it
	c.Status(200)
}