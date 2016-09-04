package controllers

import (
    "github.com/astaxie/beego"
    "github.com/lenfree/barcode/models"
)

// Status about v1 app
type StatusController struct {
    beego.Controller
}

// @Title Get
// @Description app status
// @Success 200 {object} models.AppStatus
// @Failure 500 Internal server error
// @router / [get]
func (v *StatusController) Get() {
    status := models.Status()
    v.Data["json"] = &status
    v.ServeJSON()
}
