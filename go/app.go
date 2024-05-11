package main


import (
	"fmt"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this a demonstration of how this Go app works")
	fmt.Fprintln(w, "App is working...")
	fmt.Fprintln(w, "App finished working...")
	fmt.Fprintln(w, "Goodbye, World!")
}



func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
