package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	port  = flag.String("port", "11130", "Specifiy Port(default:11130)")
	https = flag.Bool("https", false, "Enable or disable https(default:false)")
)

func Hello(w http.ResponseWriter, r *http.Request) {
	message := "Hello Hepsiburada from " + string(r.URL.Query().Get("Name"))
	io.WriteString(w, message)
	return

}

func main() {
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	fmt.Println("Serving on:" + *port)
	if *https == true {
		http.ListenAndServeTLS(":"+*port, "./certs/server.crt", "./certs/server.key", r)
	} else {
		http.ListenAndServe(":"+*port, r)
	}

}
