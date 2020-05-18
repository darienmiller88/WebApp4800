package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct{
	db *sql.DB
}

//In my database, the users that are inserted are meant to be applicants in my bank account website

//Method to connect to a mysql database
func (m *MySQL) connectToDB(){
	dBConnection, err := sql.Open("mysql", "root:nintendowiiu000@/webapp")
	if err != nil {
		fmt.Println("Connection Failed!!")
		return
	}

	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
		return
	}

	fmt.Println("CONNECTED TO MYSQL WOOO")
	m.db = dBConnection
}

func (m *MySQL) getEmail(userName string) string {
	var emailColumn string 
	emailResult := "SELECT email FROM users WHERE userName = ?"

	numRowsForEmail := m.db.QueryRow(emailResult, userName)
	emailErr := numRowsForEmail.Scan(&emailColumn)
	
	if emailErr != nil {
		if emailErr == sql.ErrNoRows{
			return ""
		}
		panic(emailErr)
	}

	return emailColumn
}

func (m *MySQL) checkUserNameAndPassword(userName string, password string) bool{
	var userNameColumn, passwordColumn string 
	checkUserName := "SELECT userName FROM users WHERE userName = ?"
	checkPassword := "SELECT password FROM users WHERE password = ?"
	
	//Query for username
	numRowsForUserName := m.db.QueryRow(checkUserName, userName)
	userNameErr := numRowsForUserName.Scan(&userNameColumn)

	//and then password
	numRowsForPassword := m.db.QueryRow(checkPassword, password)
	passWordErr := numRowsForPassword.Scan(&passwordColumn)

	if userNameErr != nil || passWordErr != nil {
		if userNameErr == sql.ErrNoRows || passWordErr == sql.ErrNoRows{

			return false
		}
		panic(userNameErr)
	}

	return true
} 

func (m *MySQL) getUserByID(id int) bool{
	// user, err := m.db.Query("SELECT * FROM users WHERE id=?", id)

	// if err != nil {
	// 	panic(err.Error())
	// }
	
	// fmt.Println(user)
	var column string 
	sqlStatement := "SELECT firstName FROM users WHERE id = ?"
	row := m.db.QueryRow(sqlStatement, id)
	err := row.Scan(&column)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}

	fmt.Println(column)
	return true
}

//CREATE - This function will take data from the front end, and add it into the database
func(m *MySQL) insertIntoDB(firstName string, lastName string, password string, email string, userName string){
	insertResult, err := m.db.Prepare("INSERT INTO users(firstName, lastName, password, email, userName) VALUES(?,?,?,?,?)")
	
	if err != nil {
		panic(err.Error())
	}

	result, _ := insertResult.Exec(firstName, lastName, password, email, userName)
	fmt.Println(result)
}

//DELETE - will allow me to delete a user based on their ID.
func (m *MySQL) deleteRowByID(id int){
	if(m.getUserByID(id)){
		deleteResult, err := m.db.Prepare("DELETE FROM users WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		deleteResult.Exec(id)
	}else{
		fmt.Printf("Zero rows found for id: %d", id)
	}

}

func (m *MySQL) closeDB(){
	m.db.Close()
}
