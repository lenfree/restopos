package test

import (
    "encoding/json"
    "testing"
    _ "github.com/lenfree/barcode/routers"
    "io/ioutil"
    "strings"

    . "github.com/smartystreets/goconvey/convey"
    "github.com/nicholasf/fakepoint"
    "github.com/lenfree/barcode/models"
)

func TestPostCategory(t *testing.T) {
    data, _ := json.Marshal(models.Category{ID: 1, Name: "Drinks"})

    maker := fakepoint.NewFakepointMaker()
    fakepoint := maker.NewPost("http://localhost:8080/v1/category", 200)
    fakepoint.SetResponse(string(data)).SetHeader("Content-Type", "application/json")

    reqData := `{"category_name": "Drinks"}`
    reqBody := strings.NewReader(reqData)

    resp, _ := maker.Client().Post("http://localhost:8080/v1/category", "application/json", reqBody)
    body, _ := ioutil.ReadAll(resp.Body)

    Convey("Subject: Create category\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(resp.StatusCode, ShouldEqual, 200)
            })
            Convey("Status Code return ID:1 and Name: Drinks", func() {
                    So(string(body), ShouldEqual, string(data))
            })
    })
}

func TestGetCategory(t *testing.T) {
    data, _ := json.Marshal([]models.Category{
            {ID: 1, Name: "App"},
            {ID: 2, Name: "Drinks"},
    })

    maker := fakepoint.NewFakepointMaker()
    fakepoint := maker.NewGet("http://localhost:8080/v1/category", 200)
    fakepoint.SetResponse(string(data)).SetHeader("Content-Type", "application/json")

    resp, _ := maker.Client().Get("http://localhost:8080/v1/category")
    body, _ := ioutil.ReadAll(resp.Body)

    var response []*models.Category
    _ = json.Unmarshal(body, &response)

    Convey("Subject: Create category\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(resp.StatusCode, ShouldEqual, 200)
            })
            Convey("should return category all category", func() {
                    So(string(body), ShouldEqual, string(data))
            })
            Convey("should category count equal 2", func() {
                    So(len(response), ShouldEqual, 2)
            })
    })
}

func TestGetCategoryID(t *testing.T) {
    data, _ := json.Marshal([]models.Category{
            {ID: 2, Name: "App"},
    })

    maker := fakepoint.NewFakepointMaker()
    fakepoint := maker.NewGet("http://localhost:8080/v1/category/1", 200)
    fakepoint.SetResponse(string(data)).SetHeader("Content-Type", "application/json")

    resp, _ := maker.Client().Get("http://localhost:8080/v1/category/1")
    body, _ := ioutil.ReadAll(resp.Body)

    var response []*models.Category
    _ = json.Unmarshal(body, &response)

    Convey("Subject: Create category\n", t, func() {
            Convey("Status Code Should Be 200", func() {
                    So(resp.StatusCode, ShouldEqual, 200)
            })
            Convey("should return category all category", func() {
                    So(string(body), ShouldEqual, string(data))
            })
            Convey("should category count equal 2", func() {
                    So(len(response), ShouldEqual, 1)
            })
    })
}
