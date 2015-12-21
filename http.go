package main

import (
	"fmt"
	"net/http"
	"strconv"

	"./fizzbuzz"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := r.FormValue("n")
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		list := fizzbuzz.FizzBuzz(i)
		for _, v := range list {
			fmt.Fprintf(w, fmt.Sprintf("%v\n", v))
		}
	})
	http.ListenAndServe(":8080", nil)
}
