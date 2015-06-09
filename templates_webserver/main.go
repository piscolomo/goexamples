package main

import (
  "html/template"
  "net/http"
)

type User struct {
  FirstName string
  LastName  string
  Edad      int
}

var users []User
var templates = template.Must(template.ParseGlob("templates/*"))

func TodoMainHandler(res http.ResponseWriter, req *http.Request){
  err := templates.ExecuteTemplate(res, "base", users)
  if err != nil {
    http.Error(res, err.Error(), http.StatusInternalServerError)
    return
  }
}

func main() {
  users = make([]User, 0)
  users = append(users, User{ "Tyrion", "Lannister", 37})
  users = append(users, User{ "Daenerys", "Targaryen", 20})
  users = append(users, User{ "Jon", "Snow", 20})
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  http.HandleFunc("/", TodoMainHandler)
  http.ListenAndServe(":8089", nil)
}
