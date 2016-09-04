package test

import (
    "fmt"
    "runtime"
    "path/filepath"
    _ "github.com/lenfree/barcode/routers"
    "os"

    "github.com/astaxie/beego"
    _ "github.com/mattn/go-sqlite3"
    "github.com/astaxie/beego/orm"
)

func init() {
    _, file, _, _ := runtime.Caller(1)
    apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
    beego.TestBeegoInit(apppath)

    orm.RegisterDriver("sqlite", orm.DRSqlite)
    orm.RegisterDataBase("default", "sqlite3", "my_test.db")

    name := "default"
    force := false
    verbose := false
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        fmt.Println(err)
    }
}

func tearDown() {
    beego.Debug("Starting teardown")
    _, file, _, _ := runtime.Caller(1)
    apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
    db := apppath + string(filepath.Separator) + "my_test.db"
    os.Remove(db)
}
