package main

import (
	"fmt"
	"net/http"
	"log"
)

const PORT = ":80"

func main(){
	fmt.Println("The server is running")
	
	http.HandleFunc("/", sendIndex )

	log.Fatal(http.ListenAndServe(PORT, nil))
}

func sendIndex(w http.ResponseWriter, r *http.Request){
	fs := http.FileServer(http.Dir("tmp"))
	if r.URL.Path[len("/"):] == "style.css" {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	}
	fs.ServeHTTP(w,r)
}
