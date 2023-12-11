# web-server

`main.go` starts a basic web server that simulates a store's API. 
- The database is simulated by a dictionary that maps item names to prices.
- The API supports calls to list items, add, update, read an item, and delete an item.
- The database is initialized with two entries: "shoes": 50 and "socks": 5.
- There are basic error checks to prevent in each handler. 

List

```shell
$ curl localhost:8080/list
shoes: $50.00
socks: $5.00
```

Add
```shell
$ curl localhost:8080/add?item=ties\&price=A
invalid price: "A"

$ curl localhost:8080/add?item=ties\&price=13
added ties with price $13.00

$ curl localhost:8080/list
shoes: $50.00
socks: $5.00
ties: $13.00

$ curl localhost:8080/add?item=tie\&price=50
duplicate item: "tie"
```

Update
```shell
$ curl localhost:8080/update?item=laces\&price=14
item not found: "laces"

$ curl localhost:8080/update?item=tie\&price=A
invalid price: "A"

$ curl localhost:8080/update?item=tie\&price=3
updated laces with price $3.00
```

Read
```shell
$ curl localhost:8080/read?item=gloves
item not found "gloves"

$ curl localhost:8080/read?item=tie
item laces has price $3.00
```

Delete
```shell
$ curl localhost:8080/delete?item=gloves
item not found "gloves"

$ curl localhost:8080/delete?item=tie
deleted tie

$ curl localhost:8080/list
shoes: $50.00
socks: $3.00
```
