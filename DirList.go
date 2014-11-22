package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var cwd string
var err error

func main() {
	var port = flag.String("p", ":8282", "Port")
	flag.Parse()

	cwd, err = os.Getwd()
	if err != nil {
		log.Println(err.Error())
		return
	}
	http.Handle("/", http.FileServer(http.Dir(cwd)))
	log.Println("Serving", cwd)
	http.ListenAndServe(*port, Logger(http.DefaultServeMux))
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
