package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/monaco", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "monaco.html")
	})
	http.HandleFunc("/ace", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ace.html")
	})

	http.HandleFunc("/runcode", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Printf("%s\n", dump)
		fmt.Fprintf(w, "%s", dump)
	})

	fmt.Println("Listening at :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
