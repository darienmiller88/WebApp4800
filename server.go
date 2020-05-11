package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	// "encoding/json"
	// "math/rand"
	// "strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type FireEmblemCharacter struct {
}

type Waifu struct {
}

func main() {
	//create the router that will be used to create our routes and methods
	goRouter := mux.NewRouter()

	//This is the golang way of serving static files lmao
	goRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//routes for "api"
	goRouter.HandleFunc("/", landingPage).Methods("GET")
	goRouter.HandleFunc("/FireEmblemCharacters", getFireEmblemCharacters).Methods("GET")
	goRouter.HandleFunc("/{character}", getFireEmblemCharacter).Methods("GET")
	goRouter.HandleFunc("/addCharacter", addCharacter).Methods("POST")
	goRouter.HandleFunc("/addWaifu", addWaifu).Methods("PUT")
	goRouter.HandleFunc("/deleteCharacter", deleteCharacter).Methods("DELETE")

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(getPort("8080"), goRouter))
}

func landingPage(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, request, "templates/landingPage.html")
}

//GET request
func getFireEmblemCharacters(response http.ResponseWriter, request *http.Request) {
	// fmt.Fprint(response, "here are the characters!")
	renderTemplate(response, request, "templates/feCharacters.html")

}

//GET request
func getFireEmblemCharacter(response http.ResponseWriter, request *http.Request) {
	request.FormValue("name")
}

//POST request
func addCharacter(response http.ResponseWriter, request *http.Request) {

}

//DELETE request
func deleteCharacter(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "delete called")
}

//PUT request
func addWaifu(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "waifu called")
	renderTemplate(response, request, "templates/waifu.html")
}

//Function to render an HTML file to the client side
func renderTemplate(response http.ResponseWriter, request *http.Request, fileName string) {
	//first, we
	myTemplate := template.Must(template.ParseFiles(fileName))
	if err := myTemplate.Execute(response, nil); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
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
