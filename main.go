package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		}
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
