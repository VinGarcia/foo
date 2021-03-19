package foo

import (
	"fmt"
	"net/http"

	"github.com/vingarcia/bar"
)

func Main() {
	port := "1358"
	fmt.Printf("Starting server on port %s...\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "VinGarcia says Oi Caio!!")
		fmt.Printf("received request: %#v\n", r)
	})

	fmt.Println(bar.Bar{})

	http.ListenAndServe(":"+port, nil)
}
