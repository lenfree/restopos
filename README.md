Restopos
=======

A resto POS that includes ordering, inventory

Status: Work In Progress

Getting started:
----------------

Contributing:
```
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
```

Install required packages:
```
$ go get github.com/gin-gonic/gin
$ go get github.com/coopernurse/gorp
$ go get github.com/mattn/go-sqlite3
$ go get github.com/mailgun/godebug
```

Start http server:
```
$ go run foods.go main.go structs.go
```

Create food entry:
```
$ curl -XPOST -H 'Content-Type: application/json' -d '{"name": "lechon", "price": 100.00
 }' http://localhost:8080/foods'
```

Should get similar JSON response such as below:
```
{"name":"lechon","prie":100,"result":"Success"}"
```


Retrieve all foods:
```
$ curl -XGET http://localhost:8080 | jq .
```

Should get similar JSON response such as below:
```
{
  "0": {
    "Id": 1,
    "Created": 1438831867970383600,
    "Name": "lechon",
    "Price": 100
  },
  "1": {
    "Id": 2,
    "Created": 1438833740828058400,
    "Name": "lechon",
    "Price": 20
  }
}
```


