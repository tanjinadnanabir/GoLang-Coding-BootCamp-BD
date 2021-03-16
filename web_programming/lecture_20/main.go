package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/mateors/mcb"
)

// var db *sql.DB
var db *mcb.DB
var err error

func init() {

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb

	// db, err = sql.Open("mysql", "root:r00tc0de@tcp(127.0.0.1:3306)/loruki_db")

	// if there is an error opening the connection, handle it

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer the close till after the main function has finished
	// executing

	// defer db.Close()

	// fmt.Println("db connection successful")

	// Couchbase connection block

	db = mcb.Connect("localhost", "tanjinadnanabir", "r00tc0de")

	res, err := db.Ping()

	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	//http.HandleFunc("/requestc", requestc)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")

	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

// func request(w http.ResponseWriter, r *http.Request) {

// method-1
// 	name := r.FormValue("name")
// 	company := r.FormValue("company")
// 	email := r.FormValue("email")

// fmt.Println(name, company, email)
// fmt.Fprintf(w, `received %s %s %s`, name, company, email) //response

// method-2
// r.ParseForm()
// for key, val := range r.Form {
// 	fmt.Println(key, val)
// }

// 	qs := "INSERT INTO `requests` (`id`, `name`, (NULL, '%s', '%s', '%s', '1');"`company`, `email`, `status`) VALUES
// 	sql := fmt.Sprintf(qs, name, company, email)
// 	fmt.Println(sql)
// 	insert, err := db.Query(sql)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Fprintf(w, `OK`)
// }

// like mysql table schema

type RequestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
}

func request(w http.ResponseWriter, r *http.Request) {

	// method-1

	// name := r.FormValue("name")
	// company := r.FormValue("company")
	// email := r.FormValue("email")

	// fmt.Println(name, company, email)

	// fmt.Fprintf(w, `received %s %s %s`, name, company, email) //response

	// method-2

	r.ParseForm()

	for key, val := range r.Form {
		fmt.Println(key, val)
	}

	var reqTable RequestTable

	r.Form.Add("bucket", "rootcode")
	r.Form.Add("aid", "requests::3") // we will update later
	r.Form.Add("type", "requests")
	r.Form.Add("status", "1")

	pRes := db.Insert(r.Form, &reqTable)

	fmt.Println(pRes.Status, pRes.Errors)

	fmt.Fprintf(w, `OK`)
}
