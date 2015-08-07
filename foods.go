package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func GetFoods(c *gin.Context) {
	var foods []Food
	_, err := dbmap.Select(&foods, "select * from foods order by food_id")
	CheckErr(err, "Select failed")
	content := gin.H{}
	for k, v := range foods {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

func FoodPost(c *gin.Context) {
	var json Food

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	food := CreateFood(json.Name, json.Price, json.CategoryId)
	if food.Name == json.Name {
		content := gin.H{
			"result": "Success",
			"id":     food.Id,
			"name":   food.Name,
			"price":  food.Price,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}

func CreateFood(name string, price float32, categoryId int64) Food {
	food := Food{
		Created:    time.Now().UnixNano(),
		Name:       name,
		Price:      price,
		CategoryId: categoryId,
	}

	err := dbmap.Insert(&food)
	CheckErr(err, "Insert failed")
	return food
}

func FoodDelete(c *gin.Context) {
	name := c.Param("name")

	count := DeleteFood(name)
	if count > 0 {
		content := gin.H{
			"result":      "Success",
			"row deleted": count,
		}
		c.JSON(200, content)
	} else {
		content := gin.H{
			"result":     "There's no food named " + name + ".",
			"row delete": 0,
		}
		c.JSON(404, content)
	}
}

func DeleteFood(name string) int64 {
	foodItem, err := GetFood(name)
	if err != nil {
		count := 0
		return int64(count)
	}

	count, err := dbmap.Delete(&foodItem)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func GetFood(name string) (Food, error) {
	var food Food
	err := dbmap.SelectOne(&food, "select * from foods where Name=?", name)
	if err != nil {
		return food, err
	}
	return food, nil
}

func FoodGet(c *gin.Context) {
	name := c.Param("name")

	foodItem, err := GetFood(name)

	if err != nil {
		content := gin.H{
			"result": "No match found",
		}
		c.JSON(404, content)
	} else {
		content := gin.H{
			"result": "Success",
			"name":   foodItem.Name,
			"price":  foodItem.Price,
		}
		c.JSON(200, content)
	}
}

func FoodPatch(c *gin.Context) {
	name := c.Param("name")
	var json Food

	c.Bind(&json)
	count := UpdateFood(name, json.Name, json.Price)
	if count > 0 {
		content := gin.H{
			"result":      "Success",
			"row updated": count,
		}
		c.JSON(200, content)
	} else {
		content := gin.H{
			"result":      "No food named " + name + " found.",
			"row updated": 0,
		}
		c.JSON(404, content)
	}
}

func UpdateFood(currentName, newName string, newPrice float32) int64 {
	foodItem, err := GetFood(currentName)
	if err != nil {
		count := 0
		return int64(count)
	}

	foodItem.Name = newName
	foodItem.Price = newPrice

	count, err := dbmap.Update(&foodItem)
	if err != nil {
		count := 0
		return int64(count)
	}
	return count
}
