package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//handler
// * means I'm sending a reference not a copy
// the code will not see if the data is changed but if we pass in a reference/pointer, it will
// in Go, any type can be passed in as referenced OR copy
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // r = request

	//header must be set befoe body
	w.Header().Add("Content-Type", "text/plain") // key (header name),value (header value)
	w.Write([]byte("hello " + name + "\n"))

	// w.Write("hello world!") --error
	w.Write([]byte("hello world!"))

}

// where program starts
func main() {
	// var addr string = os.Getenv("ADDR")
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		log.Fatal("please set ADDR environment vaiable")
	}

	http.HandleFunc("/hello", helloHandler) // when someone does GET/POST/.. on /hello, you'll pass in the pointer to the function

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil)) // nil == null

}
