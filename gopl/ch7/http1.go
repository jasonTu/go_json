package main

import (
	"log"
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

/*
package http

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request
}
func ListenAndServe(address string, h Handler) error
 */

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range(db) {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":8000", db))
}
