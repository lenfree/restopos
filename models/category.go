package models

import (
    "errors"
    "fmt"
    "github.com/astaxie/beego/validation"
    _ "github.com/mattn/go-sqlite3"
    "github.com/astaxie/beego/orm"
)

type Category struct {
        ID          int `json:"category_id"`
        Name        string `json:"category_name"`
}

type Product struct {
        ID          int `json:"product_id"`
        Name        string `json:"product_name"`
        Price       float64 `json:"product_price"`
        Category    *Category   `orm:"rel(one)"`
}

func AddCategory(c Category) (category Category, err error) {
    o := orm.NewOrm()
    valid := validation.Validation{}
    valid.Required(c.Name, "name")
    if valid.HasErrors() {
        return category, errors.New("hshshshs")
    }
    id, err := o.Insert(&c)
    if err == nil {
        fmt.Println(id)
    } else {
        fmt.Println(err)
    }
    return c, nil
}

func GetCategory(c string) (string, error) {
    return "hello", nil
}
