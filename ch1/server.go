package ch1

import (
	"log"
	"net/http"
	"strconv"
)

func InitGIFServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	cycles, err := strconv.ParseFloat(params.Get("cycles"), 64)

	if err != nil {
		w.Write([]byte("Invalid cycle value"))
		return
	}

	Lissajous(w, cycles)
}
