package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


//In my database, the users that are inserted are meant to be applicants in my bank account website

//Method to connect to a mysql database
func connectToDB() *sql.DB{
	dBConnection, err := sql.Open("mysql", "root:password@/webapp")
	if err != nil {
		log.Fatal("Connection Failed!!")
	}

	//Ping the database to make sure the connection is stable.
	err = dBConnection.Ping()
	if err != nil {
		log.Fatal("Ping Failed!!")
	}

	return dBConnection
}

//CREATE - This function will take data from the front end, and add it into the database
func insertIntoDB(user User){
	db := connectToDB()
	insertIntoSQL, insertErr := db.Prepare("INSERT INTO users(firstName, lastName, password, email, userName) VALUES(?,?,?,?,?)")
	errorHandleSQLQuery(insertErr)
	
	result, resultErr := insertIntoSQL.Exec(user.FirstName, user.LastName, user.Password, user.Email, user.Username)
	errorHandleSQLQuery(resultErr)

	db.Close()
	fmt.Println(result)
}

//READ - This function will return a row with the given information for a user with "userName"
func getUserByUserName(userName string) User {
	db := connectToDB()

	var user User 
	userSQLQuery := "SELECT * FROM users WHERE userName = ?"

	//Query the row in the table to return a row with the following userName
	var numRowsForUser *sql.Row = db.QueryRow(userSQLQuery, userName)
	var userErr error = numRowsForUser.Scan(&user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Username)
	
	if userErr != nil {
		//If there is no user for the username the user entered, return an empty User object
		if userErr == sql.ErrNoRows{
			return User{}
		}
		panic(userErr)
	}

	db.Close()
	return user
}

//UPDATE - This function will update the old password with a new password for the user
func updateUserPassword(newPassWord string, username string){
	db := connectToDB()

	updateSQLQuery, updateErr := db.Prepare("UPDATE users SET password = ? WHERE userName = ?")
	errorHandleSQLQuery(updateErr)

	result, resultErr := updateSQLQuery.Exec(newPassWord, username)
	errorHandleSQLQuery(resultErr)

	fmt.Println(result)
	db.Close()
}

//DELETE - will allow me to delete a user based on their username.
func deleteRowByUserName(userName string){
	db := connectToDB()
	
	if(getUserByUserName(userName).Username == ""){
		fmt.Printf("Zero rows found for userName: %s", userName)
	}else{
		deleteSQLQuery, deleteErr := db.Prepare("DELETE FROM users WHERE id=?")
		errorHandleSQLQuery(deleteErr)

		result, resultErr := deleteSQLQuery.Exec(userName)
		errorHandleSQLQuery(resultErr)

		fmt.Println(result)
	}
	db.Close()
}

func errorHandleSQLQuery(err error) {
	if err != nil {
		panic(err.Error())
	}
}
