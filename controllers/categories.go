package controllers

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "github.com/lenfree/barcode/models"
)

// Category Object
type CategoryController struct {
    beego.Controller
}

// @Title GetAll
// @Description get all Food Categories
// @Success 200 {object} models.Category
// @router / [get]
func (c *CategoryController) GetAll() {
    category, err := models.GetCategory()
    if err != nil {
        c.Data["json"] = err.Error()
    } else {
        c.Data["json"] = category
    }
    c.ServeJSON()
}

// @Title Get
// @Description get food category by id
// @Param    uid        path     string    true        "The uid of food category"
// @Success 200 {string} models.Category.Name
// @Failure 403 uid is empty
// @router /:uid [get]
func (c *CategoryController) Get() {
    uid := c.GetString(":uid")
    category, err := models.GetCategoryByID(uid)
    if err != nil {
        c.Data["json"] = err.Error()
    } else {
        c.Data["json"] = category
    }
    c.ServeJSON()
}

// @Title CreateCategory
// @Description create categories
// @Param    body        body    models.Category    true        "body for category content"
// @Success 200 {int} models.Category.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CategoryController) Post() {
    var category models.Category
    json.Unmarshal(c.Ctx.Input.RequestBody, &category)
    result, err := models.AddCategory(category)
    if err != nil {
        c.Abort("400")
    } else {
        c.Data["json"] = result
    }
    c.ServeJSON()
}
