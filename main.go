package main

import (
	"github.com/Makinde1034/email-verifier/routes"
	"net/http"
)


func main(){
	mux := routes.RegisterRoutes();
	http.ListenAndServe(":5000",mux);
}
