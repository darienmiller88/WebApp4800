package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

type Routes struct {
	MyRouter *mux.Router
	UserData User
	db MySQL
}

type User struct{
	FirstName, LastName, Password, Email, Username, Welcome string
	//Done bool
}

//Method to assign a new mux router to an instance of a "Routes" object and
func (r *Routes) createRoute() {
	r.MyRouter = mux.NewRouter()
	r.db.connectToDB()
	// r.UserData = User{
	// 	Name: "Anonymous", 
	// 	Country: "USA",
	// 	Welcome: "Landing",
	// 	Done: false,
	// }

	//This is the golang way of serving static files lmao
	r.MyRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//routes for web app
	r.MyRouter.HandleFunc("/", r.landingPage).Methods("GET")
	r.MyRouter.HandleFunc("/login", r.verifyLogIn).Methods("POST")
	r.MyRouter.HandleFunc("/login", r.loginPage).Methods("GET")
	r.MyRouter.HandleFunc("/signup", r.signUp).Methods("GET")
	r.MyRouter.HandleFunc("/aboutus", r.aboutUs).Methods("GET")
	r.MyRouter.HandleFunc("/", r.deleteStuff).Methods("DELETE")

}

//GET request
func (r *Routes) landingPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("landing page called")
	r.renderTemplate(writer, request, "templates/landingPage.html", r.UserData)
}

func (r *Routes) loginPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("login page called")
	r.renderTemplate(writer, request, "templates/login.html", r.UserData)
}

//POST request
func (r *Routes) verifyLogIn(writer http.ResponseWriter, request *http.Request){
	fmt.Println("POST request called!")
	// Call ParseForm() to parse the raw query and update request.PostForm and request.Form.
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	fmt.Printf("Post from website! request.PostFrom = %v\n", request.PostForm)
	r.UserData.FirstName = request.PostForm.Get("firstName")
	r.UserData.LastName = request.PostForm.Get("lastName")
	r.UserData.Password = request.PostForm.Get("password")
	r.UserData.Email = request.PostForm.Get("email")

	r.db.insertIntoDB(r.UserData.FirstName, r.UserData.LastName, r.UserData.Password, r.UserData.Email)
	r.redirectHome(writer, request)
	// name := request.FormValue("name")
	// address := request.FormValue("address")
	// fmt.Fprintf(writer, "Name = %s\n", name)
	// fmt.Fprintf(writer, "Address = %s\n", address)
}

//GET request
func (r *Routes) signUp(writer http.ResponseWriter, request *http.Request){
	fmt.Println("sign up page called")
	r.renderTemplate(writer, request, "templates/signup.html", nil)
}

//GET request
func (r *Routes) aboutUs(writer http.ResponseWriter, request *http.Request){
	fmt.Println("about us page called")
	r.renderTemplate(writer, request, "templates/aboutUs.html", nil)
}

func (r *Routes) redirectHome(writer http.ResponseWriter, request *http.Request) {
    http.Redirect(writer, request, "/", 301)
}

func (r *Routes) deleteStuff(writer http.ResponseWriter, request *http.Request){
	fmt.Println("Delete called!")
}

func (r *Routes) renderTemplate(response http.ResponseWriter, request *http.Request, fileName string, data interface{}){
	myTemplate := template.Must(template.ParseFiles(fileName))
	if err := myTemplate.Execute(response, data); err != nil {
		fmt.Println("error :(")
		//http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
