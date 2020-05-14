package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	// "encoding/json"
	// "math/rand"
	// "strconv"
)

type FireEmblemCharacter struct {
	name string
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
	goRouter.HandleFunc("/myapi/FireEmblemCharacters", getFireEmblemCharacters).Methods("GET")
	goRouter.HandleFunc("/myapi/{character}", getFireEmblemCharacter).Methods("GET")
	goRouter.HandleFunc("/myapi/addCharacter", addCharacter).Methods("POST")
	goRouter.HandleFunc("/myapi/addWaifu", addWaifu).Methods("GET")
	goRouter.HandleFunc("/myapi/deleteCharacter", deleteCharacter).Methods("DELETE")

	fmt.Println("Server is running!")
	log.Fatal(http.ListenAndServe(getPort("3000"), goRouter))
}

func landingPage(writer http.ResponseWriter, request *http.Request) {
	// request.FormValue("name")
	//fmt.Fprintf(writer, "landing")
	renderTemplate(writer, request, "templates/landingPage.html")
}

//GET request
func getFireEmblemCharacters(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, request, "templates/feCharacters.html")
}

//GET request
func getFireEmblemCharacter(writer http.ResponseWriter, request *http.Request) {

}

//POST request
func addCharacter(writer http.ResponseWriter, request *http.Request) {

}

//DELETE request
func deleteCharacter(writer http.ResponseWriter, request *http.Request) {

}

//PUT request
func addWaifu(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "waifu called")
	renderTemplate(writer, request, "templates/waifu.html")
}

func renderTemplate(response http.ResponseWriter, request *http.Request, fileName string) {
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
