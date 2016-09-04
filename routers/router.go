// @APIVersion 1.0.0
// @Title Barcode API
// @Description Barcode
// @Contact lenfree.yeung@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
        "github.com/lenfree/barcode/controllers"
        "github.com/astaxie/beego"
)

func init() {
        ns := beego.NewNamespace("/v1",
                beego.NSNamespace("/user",
                        beego.NSInclude(
                                &controllers.UserController{},
                        ),
                ),
                beego.NSNamespace("/status",
                        beego.NSInclude(
                                &controllers.StatusController{},
                        ),
                ),
                beego.NSNamespace("/category",
                        beego.NSInclude(
                                &controllers.CategoryController{},
                        ),
                ),
        )
        beego.AddNamespace(ns)
}
