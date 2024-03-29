# xkcd-search

`xkcd-load.go` performs an http get request to xkcd.com to pull comic data from its JSON database until it gets two 404 responses (since comic #404 is missing). File is > 2mb so the load takes a few minutes.

```shell
$ go run ./load/xkcd-load.go xkcd2.json
read 2864 comics
```

`xkcd-find.go` takes sevaeral command line arguments; the xkcd.json file generated by xkcd-load.go, followed by a series of search terms. The xkcd.json file is opened, the json in the file is decoded into and loaded into memory. We then loop over each item and check if the search terms are contained in the title or transcript. 

```shell
$ go run ./find/xkcd-find.go ./xkcd.json flinstone
read 2864 comics
https://xkcd.com/1491/ 2/25/2015 Stories of the Past and Future
found 1 comics
```
