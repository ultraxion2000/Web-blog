package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name                  string `json:"name"`
	Age                   uint16 `json:"age"`
	Projects              int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has projects "+
		" equal: %d", u.Name, u.Age, u.Projects)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	client := User{"Alexander", 23, 19, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, client)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var bob User = ...
	//bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	//bob := User{"Bob", 25, -50, 4.2, 0.8}

	//handleRequest()

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goland")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Установка данных

	// insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('alex', 23)")
	// if err != nil {
	// 	panic(err)
	// }

	// defer insert.Close()

	// Выборка данных
	res, err := db.Query(("SELECT `name`, `age` FROM `users`"))
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	}

}
