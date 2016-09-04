package main

import (
    _ "github.com/lenfree/barcode/routers"
    "github.com/astaxie/beego"
)

func main() {
    if beego.BConfig.RunMode == "dev" {
        beego.BConfig.WebConfig.DirectoryIndex = true
            beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
    }
    beego.Run()
}
