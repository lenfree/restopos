package test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    _ "github.com/lenfree/barcode/routers"

    "github.com/astaxie/beego"
    . "github.com/smartystreets/goconvey/convey"
)

func TestGetStatus(t *testing.T) {
    r, _ := http.NewRequest("GET", "/v1/status", nil)
    w := httptest.NewRecorder()
    beego.BeeApp.Handlers.ServeHTTP(w, r)

    beego.Trace("testing", "TestGetStatus", "Code[%d]\n%s", w.Code, w.Body.String())

    Convey("Subject: App status\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(w.Code, ShouldEqual, 200)
            })
    })
}
