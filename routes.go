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

//A user in my web app will have the following information, typical of a bank customer
type User struct{
	FirstName, LastName, Password, Email, Username, Welcome string
	UserAccounts []Account
	LoginSuccessful bool
}

type Account struct{
	accountType, accountName string
	accountBalance float64
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
	r.MyRouter.HandleFunc("/signup", r.verifySignUp).Methods("POST")
	r.MyRouter.HandleFunc("/aboutus", r.aboutUs).Methods("GET")
	r.MyRouter.HandleFunc("/viewprofile", r.viewProfile).Methods("GET")
	r.MyRouter.HandleFunc("/viewprofile/delete", r.deleteStuff).Methods("DELETE")

}

//GET request
func (r *Routes) landingPage(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("landing page called")
	r.renderTemplate(writer, request, "templates/landingPage.html", r.UserData)
}

func (r *Routes) viewProfile(writer http.ResponseWriter, request *http.Request){
	fmt.Println("view profile")
	r.renderTemplate(writer, request, "templates/viewProfile.html", r.UserData)
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
	tempUserName := request.PostForm.Get("username")
	tempPassword := request.PostForm.Get("password")

	//If the user entered the wrong user name OR password, redirect them back to the login page
	if(!r.db.checkUserNameAndPassword(tempUserName, tempPassword)){
		http.Redirect(writer, request, "/login", 301)
	}else{
		r.UserData.Username = tempUserName
		r.UserData.Password = tempPassword
		r.UserData.Email = r.db.getEmail(tempUserName)
		http.Redirect(writer, request, "/viewprofile", 301)
	}
}

//GET request
func (r *Routes) signUp(writer http.ResponseWriter, request *http.Request){
	fmt.Println("sign up page called")
	r.renderTemplate(writer, request, "templates/signup.html", nil)
}

func (r *Routes) verifySignUp(writer http.ResponseWriter, request *http.Request){
	fmt.Println("POST request called! verifySignUp()")

	// Call ParseForm() to parse the raw query and update request.PostForm and request.Form.
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}

	//obtain the users informatino from signing up...
	fmt.Printf("Post from website! request.PostFrom = %v\n", request.PostForm)
	r.UserData.FirstName = request.PostForm.Get("firstname")
	r.UserData.LastName = request.PostForm.Get("lastname")
	r.UserData.Password = request.PostForm.Get("password")
	r.UserData.Email = request.PostForm.Get("email")
	r.UserData.Username = request.PostForm.Get("username")

	//Insert it into our database...
	r.db.insertIntoDB(r.UserData.FirstName, r.UserData.LastName, r.UserData.Password, r.UserData.Email, r.UserData.Username)

	//And redirect the user to the "view profile" where they can see their information
	http.Redirect(writer, request, "/viewprofile", 301)
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
	fmt.Fprintf(writer, "We gonna delete stuff here!")
	fmt.Println("Delete called!")
}

func (r *Routes) renderTemplate(response http.ResponseWriter, request *http.Request, fileName string, data interface{}){
	myTemplate := template.Must(template.ParseFiles(fileName))
	if err := myTemplate.Execute(response, data); err != nil {
		fmt.Println("error :(")
		//http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}
