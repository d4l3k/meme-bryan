package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for {
			select {
			case <-r.Context().Done():
				return
			default:
			}
			if _, err := w.Write([]byte("Bryan\n")); err != nil {
				log.Println(err)
				return
			}
			w.(http.Flusher).Flush()
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
