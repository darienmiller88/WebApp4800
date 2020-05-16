package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Routes struct {
	myRouter *mux.Router
}

//The structs that I am using must start with a capital
type Data struct {
	Welcome string
	Names   []string
}

//Method to assign a new mux router to an instance of a "Routes" object and
func (r *Routes) createRoute() {
	r.myRouter = mux.NewRouter()

	//This is the golang way of serving static files lmao
	r.myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//routes for web app
	r.myRouter.HandleFunc("/", landingPage).Methods("GET")
	r.myRouter.HandleFunc("/signup", signUp).Methods("GET")
	r.myRouter.HandleFunc("/aboutus", aboutUs).Methods("GET")
}

func landingPage(writer http.ResponseWriter, request *http.Request) {
	// request.FormValue("name")
	//fmt.Fprintf(writer, "landing")
	fmt.Println("landing page called")
	renderTemplate(writer, request, "templates/landingPage.html", Data{
		Welcome: "Landing",
		Names: []string{
			"Darien",
			"Denise",
			"Dalton",
			"Derick",
			"Doflamingo",
		},
	})
}

//GET request
func signUp(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("sign up page called")
	renderTemplate(writer, request, "templates/signup.html", nil)
}

//GET request
func aboutUs(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("about us page called")
	renderTemplate(writer, request, "templates/aboutUs.html", nil)
}

func renderTemplate(response http.ResponseWriter, request *http.Request, fileName string, data interface{}) {
	myTemplate := template.Must(template.ParseFiles(fileName))
	if err := myTemplate.Execute(response, data); err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
