package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

	c.Bind(&json)
	category := CreateCategory(json.Name)
	if category.Name == json.Name {
		content := gin.H{
			"result": "Success",
			"id":     category.Id,
			"name":   category.Name,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}

func CreateCategory(name string) Category {
	category := Category{
		Name: name,
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
	err := dbmap.SelectOne(&category, "select * from categories where Name=?", "ulam")
	if err != nil {
		return category, err
	}
	return category, nil
}

func CategoryGet(c *gin.Context) {
	name := c.Param("name")

	CategoryItem, err := GetCategory(name)
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

func CategoryPatch(c *gin.Context) {
	name := c.Param("name")
	var json Category

	c.Bind(&json)
	count := UpdateCategory(name, json.Name)
	if count > 0 {
		content := gin.H{
			"result":      "Success",
			"row updated": count,
		}
		c.JSON(200, content)
	} else {
		content := gin.H{
			"result":      "No category named " + name + " found.",
			"row updated": 0,
		}
		c.JSON(404, content)
	}
}

func UpdateCategory(currentName, newName string) int64 {
	CategoryItem, err := GetCategory(currentName)
	if err != nil {
		count := 0
		return int64(count)
	}

	CategoryItem.Name = newName

	count, err := dbmap.Update(&CategoryItem)
	if err != nil {
		count := 0
		return int64(count)
	}
	return count
}
