Barcode
=======

A POS refactor from PHP.

Getting started:
----------------

Contribute to repo:
```
$ git clone 
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
