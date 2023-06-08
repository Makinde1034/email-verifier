package main

import (
	"github.com/Makinde1034/email-verifier/routes"
	"net/http"
)


func main(){
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("domain,hasMX,hasSPX,sprRecord,hasDMARC,dmarcRecord\n")

	// for scanner.Scan(){
	// 	verifyDomain(scanner.Text())
	// }

	// if err:= scanner.Err();err !=nil{
	// 	log.Printf("Error: occured while scanning: %v\n",err)
	// }
	mux := routes.RegisterRoutes();
	http.ListenAndServe(":5000",mux);

	
}
