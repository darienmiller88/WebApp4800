package main

import (
	"database/sql"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct{
	db *sql.DB
}

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
func(m *MySQL) insertIntoDB(firstName string, lastName string, password string, email string){
	insertResult, err := m.db.Prepare("INSERT INTO users(firstName, lastName, password, email) VALUES(?,?,?,?)")
	
	if err != nil {
		panic(err.Error())
	}

	insertResult.Exec(firstName, lastName, password, email)
}

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
