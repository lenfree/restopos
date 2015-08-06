package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func GetCategories(c *gin.Context) {
	var categories []Category
	_, err := dbmap.Select(&categories, "select * from categories order by category_id")
	CheckErr(err, "Select failed")
	content := gin.H{}
	for k, v := range categories {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

func CategoryPost(c *gin.Context) {
	var json Category

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	category := CreateCategory(json.Name)
	if category.Name == json.Name {
		content := gin.H{
			"result": "Success",
			"name":   category.Name,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}

func CreateCategory(name string) Category {
	category := Category{
		Created: time.Now().UnixNano(),
		Name:    name,
	}
	err := dbmap.Insert(&category)
	CheckErr(err, "Insert failed")
	return category
}

func CategoryDelete(c *gin.Context) {
	name := c.Param("name")

	count := DeleteCategory(name)
	if count > 0 {
		content := gin.H{
			"result":      "Success",
			"row deleted": count,
		}
		c.JSON(200, content)
	} else {
		content := gin.H{
			"result":     "No category named " + name + " found.",
			"row delete": 0,
		}
		c.JSON(404, content)
	}
}

func DeleteCategory(name string) int64 {
	CategoryItem, err := GetCategory(name)
	if err != nil {
		count := 0
		return int64(count)
	}

	count, err := dbmap.Delete(&CategoryItem)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func GetCategory(name string) (Category, error) {
	var category Category
	err := dbmap.SelectOne(&category, "select * from categories where Name=?", name)
	if err != nil {
		return category, err
	}
	return category, nil
}

func CategoryGet(c *gin.Context) {
	name := c.Param("name")

	CategoryItem, err := GetCategory(name)
	_ = "breakpoint"
	if err != nil {
		content := gin.H{
			"result": "No match found",
		}
		c.JSON(404, content)
	} else {
		content := gin.H{
			"result": "Success",
			"name":   CategoryItem.Name,
		}
		c.JSON(200, content)
	}
}
