package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	var PORT string = ":3000"

	http.HandleFunc("/", echoString)

	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/views", func(w http.ResponseWriter, r *http.Request) {
		var views string = "Счетчик: " + strconv.Itoa(counter)
		fmt.Fprintf(w, views)
	})

	log.Fatal(http.ListenAndServe(PORT, nil))

}
