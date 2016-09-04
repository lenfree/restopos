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

// @Title Get
// @Description Food Category
// @Success 200 {object} models.Category
// @Failure 500 Internal server error
// @router / [get]
func (c *CategoryController) Get() {
    cid := c.GetString(":cid")
    if cid != "" {
        category, err := models.GetCategory(cid)
        if err != nil {
            c.Data["json"] = err.Error()
        } else {
            c.Data["json"] = category
        }
    }
    c.ServeJSON()
}

// @Title CreateCategory
// @Description create categories
// @Param    body        body    models.User    true        "body for category content"
// @Success 200 {int} models.Category.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CategoryController) Post() {
    var category models.Category
    json.Unmarshal(c.Ctx.Input.RequestBody, &category)
    result, err := models.AddCategory(category)
    if err != nil {
        c.Abort("500")
    } else {
        c.Data["json"] = result
    }
    c.ServeJSON()
}
