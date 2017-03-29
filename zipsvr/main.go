package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type zip struct {
	Zip string `json:"zip"`
	City string `json:"city"`
	State string `json:"state"`
}

type zipSlice []*zip
type zipIndex map[string]zipSlice

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

// Why not *?
// interface is always passed by pointer/reference
// responsewriter is a interface not a structure
// (zi zipIndex) = receiver
// 
func (zi zipIndex) zipsForCityHandler(w http.ResponseWriter, r *http.Request) {	// must contain 2 params and must be types http.ResponseWriter and *http.Request
	// /zips/city/seattle
	_, city := path.Split(r.URL.Path)	// take a path then separate _ = /zips/city city = /seattle
	lcity := strings.ToLower(city)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// can write directly to the writer
	encoder := json.NewEncoder(w)
	
	if err:= encoder.Encode(zi[lcity]); err != null {
		http.Error(w, "error encoding json: " + err.Error(), 
			http.StatusInternalServerError) // error occurs in server, so we send out server interal error 500s
	}
	// not yet added

}

// where program starts
func main() {
	// var addr string = os.Getenv("ADDR")
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		// log.Fatal() wirtes the message to stdout 
		// and exists with a code of 1, indicating an error
		log.Fatal("please set ADDR environment vaiable")
	}

	http.HandleFunc("/hello", helloHandler) // when someone does GET/POST/.. on /hello, you'll pass in the pointer to the function

	fmt.Printf("server is listening at %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil)) // nil == null
// end of lecture 1

	f, err := os.Open("../data/zips.json")
	if err != nil {
		log.Fatal("error opening zips file: " + err.Error())
	}

	zips := make(zipSlice, 0, 43000)
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&zip); err != nil {
		log.Fatal("error decoding zips json: " + err.Error())
	}
	fmt.Printf("loaded %d zips\n", len(zips))

	zi := make(zipIndex)

	for _, z := range zips {	// foreach
	lower := strings.ToLower(z.City)
	zi[lower] = append(zi[lower],z)
	}

	fmt.Printf("there are %d zips in Seattle\n", len(zi["seattle"]))

	// http.HandleFunc("/zips/city/", zipsForCityHandler)	// not work because of receiver
	http.HandleFunc("/zips/city/", zi.zipsForCityHandler)


}
