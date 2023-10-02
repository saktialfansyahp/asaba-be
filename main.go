package main

import (
	"fmt"
	"net/http"

	"github.com/saktialfansyahp/asaba-be.git/router"
)

func main(){
	router.DefineRoutes()
	serverAddr := ":8080"
	fmt.Printf("Server is listening on %s\n", serverAddr)
	go http.ListenAndServe(serverAddr, nil)

	select {}
}