package main

import (
    _ "github.com/lenfree/barcode/routers"
    "github.com/astaxie/beego"
    "github.com/lenfree/barcode/models"
)

func init() {
    models.New("/tmp/db.sqlite3")
}

func main() {
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
            beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}
