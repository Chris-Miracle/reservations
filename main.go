package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, World!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	fmt.Println("Listening on port 4545")
	_  = http.ListenAndServe(":4545", nil)
}
