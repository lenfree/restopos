package main

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func main() {
	// Creates a gin router with default middlewares:
	// logger and recovery (crash-free) middlewares
	router := gin.Default()

	router.GET("/", index)
	router.GET("/foods", GetFoods)
	router.POST("/foods", FoodPost)
	router.DELETE("/foods/:name", FoodDelete)
	router.GET("/foods/:name", FoodGet)
	router.POST("/categories", CategoryPost)
	router.GET("/categories", GetCategories)
	router.GET("/categories/:name", CategoryGet)
	router.DELETE("/categories/:name", CategoryDelete)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}
