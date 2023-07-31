package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("oi")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle</h1>"))
}
