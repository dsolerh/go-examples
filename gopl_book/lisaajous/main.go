package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Starting Server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		if q.Has("cycles") {
			cycles, err := strconv.Atoi(q["cycles"][0])
			if err != nil {
				fmt.Fprintf(w, "%v\n", err)
			} else {
				lissajous(w, float64(cycles))
			}
		} else {
			lissajous(w, 5)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
