package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32 //don't use floats for this in real life

// implement Stringer for formatting prints
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars //simulate the database with a map

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	//look up the item to see if it already exists, we don't want to add duplicates.
	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) //send a 400
		return
	}

	//convert the string price into a 32 bit float
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) //send a 400
		return
	}

	db[item] = dollars(p) //cast the float to the dollars type

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item]) //could return a different response code to show we created something
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	//if the item doesn't exist, we can't update it.
	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("item not found: %q", item)
		http.Error(w, msg, http.StatusNotFound) //send a 404
		return
	}

	//convert the string price into a 32 bit float
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) //send a 400
		return
	}

	db[item] = dollars(p) //cast the float to the dollars type

	fmt.Fprintf(w, "updated %s with price %s\n", item, db[item]) //could return a different response code to show we created something
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("item not found %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("item not found %q", item)
		http.Error(w, msg, http.StatusNotFound) //404
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "deleted %s\n", item)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	//add routes
	http.HandleFunc("/list", db.list) //handlefunc wants a handler
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
