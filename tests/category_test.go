package test

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    _ "github.com/lenfree/barcode/routers"
    "strings"

    "github.com/astaxie/beego"
    . "github.com/smartystreets/goconvey/convey"
)

func TestGetCategory(t *testing.T) {
    r, _ := http.NewRequest("GET", "/v1/category", nil)
    w := httptest.NewRecorder()
    beego.BeeApp.Handlers.ServeHTTP(w, r)

    beego.Trace("testing", "TestGetCategory", "Code[%d]\n%s", w.Code, w.Body.String())

    Convey("Subject: Get Category\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(w.Code, ShouldEqual, 200)
            })
    })
}

func TestPostCategory(t *testing.T) {
    reqData := `{"category_name": "Drinks"}`
    reqBody := strings.NewReader(reqData)

    r, _ := http.NewRequest("POST", "/v1/category", reqBody)
    r.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    beego.BeeApp.Handlers.ServeHTTP(w, r)

    beego.Trace("testing", "TestPostCategory", "Code[%d]\n%s", w.Code, w.Body.String())

    var response struct {
        ID        int `json:"category_id"`
        Name      string `json:"category_name"`
    }

    json.Unmarshal(w.Body.Bytes(), &response)

    Convey("Subject: Create category\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(w.Code, ShouldEqual, 200)
            })
            Convey("Return last inserted ID", func() {
                    So(response.ID, ShouldEqual, 1)
            })
            Convey("Return last inserted category name", func() {
                    So(response.Name, ShouldEqual, "Drinks")
            })
    })
}
