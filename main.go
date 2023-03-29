// File Name: main
package main

import (
	"log"
	"net/http"
)

//Write middleware
func middlewareA(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//this is executed on the way down to the handler
		log.Println("Executing MiddleWare A")
		next.ServeHTTP(w,r)
		// this is executed on the way up to the client
		log.Println("Executing MiddleWare A agian")
	})
}

//create handler function
func ourHandler(w http.ResponseWriter, r *http.Request){
	
	log.Println("Execuring the handler...")
	w.Write([]byte("Carrots\n"))
}


func main(){
	mux := http.NewServeMux()
	mux.Handle("/", middlewareA(http.HandlerFunc(ourHandler)))
	
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}