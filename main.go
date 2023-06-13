package main

import (
	"net/http"

	"github.com/Makinde1034/email-verifier/routes"

	"github.com/rs/cors"
)


func main(){

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	mux := routes.RegisterRoutes();
	http.ListenAndServe(":5000",c.Handler(mux));
}
