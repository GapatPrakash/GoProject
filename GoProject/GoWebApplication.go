package main

import "net/http"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe(":8000",nil)
}

type MyHandler struct {
	http.Handler
}