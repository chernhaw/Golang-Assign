package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@/golang")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected ")

}

type User struct {
	User       string
	Password   string
	Email      string
	User_group string
}

func main() {

	router := gin.Default()

	router.POST("/users", users)
	router.POST("/userid", userid)
	// http.HandleFunc("/", foo)
	// http.HandleFunc("/users", users)
	// http.HandleFunc("user/{id}", userbyId)
	router.Run("localhost:8080")

}

// func users(w http.ResponseWriter, r *http.Request) {
func users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * from golang.user;")
	if err != nil {

		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.User, &user.Email, &user.Password, &user.User_group)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%s, %s, %s, %s", user.User, user.Email, user.Password, user.User_group)
	}
}

func userid(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	userid := r.FormValue("userid")
	if userid == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=uft-8")
	io.WriteString(w, `<form method="post"> 
			   <input type ="text" name="q">
			   <input type ="submit">
			   </form>
			   <br>`+v)
}

/*
rows, err := db.Query("SELECT * from golang.user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.User, &user.Email, &user.Password, &user.User_group)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Printf("%s, %s, %s, %s", user.User, user.Email, user.Password, user.User_group)
	}
*/
