package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: http.HandlerFunc(dumper),
	}
	err := server.ListenAndServe()
	panic(err)
}

func dumper(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, false)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Printf("%v\n", time.Now())
	fmt.Println(string(dump) + "\n")
}
