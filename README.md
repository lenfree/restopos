Restopos
=========

A resto POS that includes ordering, inventory.
This is configured with Codeship to run tests.

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
$ go get github.com/beego/bee
$ go get github.com/nicholasf/fakepoint
$ go get github.com/mattn/go-sqlite3
```

Start http server:
```
$ bee run
```

Create food category:
```
 curl http://localhost:8080/v1/category -XPOST -H 'Content-type: application/json' -d '{"category_name": "App1" }' -i
```


Should get similar JSON response such as below:
```
HTTP/1.1 200 OK
Content-Length: 50
Content-Type: application/json; charset=utf-8
Server: beegoServer:1.7.0
Date: Wed, 07 Sep 2016 09:46:23 GMT

{
  "category_id": 1,
  "category_name": "Appet"
}


Retrieve all foods:
```
curl -XGET http://localhost:8080/v1/category -H 'Content-type: application/json' | jq .
[
  {
    "category_id": 1,
    "category_name": "test1"
  },
  {
    "category_id": 2,
    "category_name": "test1"
  },
  {
    "category_id": 3,
    "category_name": "test1"
  }
]
```


Retrieve all foods:
```
curl -XGET http://localhost:8080/v1/category/2 -H 'Content-type: application/json' | jq .
{
  "category_id": 2,
  "category_name": "categoryM"
}``


Delete specific entry:
```
$ curl -XDELETE http://localhost:8080/v1/category/2
```

Should expect below output:
```

```
