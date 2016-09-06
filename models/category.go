package models

import (
    "errors"
    "fmt"
    "github.com/astaxie/beego/validation"
    _ "github.com/mattn/go-sqlite3"
    "github.com/astaxie/beego/orm"
    "strconv"
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

func AddCategory(c Category) (Category, error) {
    o := orm.NewOrm()
    valid := validation.Validation{}
    valid.Required(c.Name, "name")
    if valid.HasErrors() {
        return c, errors.New("Error")
    }
    id, err := o.Insert(&c)
    if err == nil {
        fmt.Println(id)
    } else {
        fmt.Println(err)
    }
    return c, nil
}

func GetCategory() ([]Category, error) {
    o := orm.NewOrm()
    var categories []Category
    qs := o.QueryTable("category")
    num, err := qs.All(&categories)
    if num > 0 {
        fmt.Println(err)
    }

    return categories, nil
}

func GetCategoryByID(id string) (Category, error) {
    uid, _ := strconv.Atoi(id)
    category := Category{ID: uid}
    o := orm.NewOrm()
    err := o.Read(&category)
    if err == orm.ErrNoRows {
        fmt.Println(errors.New("not"))
        return category, errors.New("404")
    } else {
        return category, nil
    }
}
