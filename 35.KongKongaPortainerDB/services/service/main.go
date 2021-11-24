package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, os.Getenv("CONTENT"))
}

func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"),nil))
}