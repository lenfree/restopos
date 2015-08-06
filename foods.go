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
			"prie":   food.Price,
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
	_ = "breakpoint"
	err := dbmap.Insert(&food)
	checkErr(err, "Insert failed")
	return food
}
