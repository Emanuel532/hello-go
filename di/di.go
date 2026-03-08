package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "eddie")
// }

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Greet(w, "John")
	})))

}
