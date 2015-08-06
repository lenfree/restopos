package main

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Food{}, "foods").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetFoods(c *gin.Context) {
	var foods []Food
	_, err := dbmap.Select(&foods, "select * from foods order by food_id")
	checkErr(err, "Select failed")
	content := gin.H{}
	for k, v := range foods {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

func FoodPost(c *gin.Context) {
	var json Food

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	food := CreateFood(json.Name, json.Price)
	if food.Name == json.Name {
		content := gin.H{
			"result": "Success",
			"name":   food.Name,
			"price":  food.Price,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}

func CreateFood(name string, price float32) Food {
	food := Food{
		Created: time.Now().UnixNano(),
		Name:    name,
		Price:   price,
	}
	err := dbmap.Insert(&food)
	checkErr(err, "Insert failed")
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
