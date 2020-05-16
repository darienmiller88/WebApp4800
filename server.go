package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "encoding/json"
	// "math/rand"
	// "strconv"
)

func main() {
	//create the router that will be used to create our routes and methods
	routeObject := Routes{}

	//Call the function that will assign a mux router
	routeObject.createRoute()

	//Finally, use the mux router that belongs to our route object to assign to the listen and serve function
	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(getPort("3000"), routeObject.myRouter))
}

func getPort(defaultPort string) string {
	var port string = os.Getenv("PORT")

	// Set a default port if there is nothing in the environment
	if port == "" {
		port = defaultPort
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
