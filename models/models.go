package models

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)

type SQLite3 orm.Ormer

func init() {
   orm.RegisterDriver("sqlite3", orm.DRSqlite)
   orm.RegisterDataBase("default", "sqlite3", "/tmp/test.sqlite")
   orm.RegisterModel(new(Category), new(Product))
   // Database alias.
   name := "default"

   // Drop table and re-create.
   force := true

   // Print log.
   verbose := true

   // Error.
   err := orm.RunSyncdb(name, force, verbose)
   if err != nil {
       fmt.Println(err)
   }
}

func New(connection string) SQLite3 {
   o := orm.NewOrm()
   o.Using("default")
   return o
}
