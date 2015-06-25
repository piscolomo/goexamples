package main

import "fmt"

type Company struct{
  name string
}

type Post struct{
  content string
}

type User struct{
  name string
  company Company
  posts []Post
}

func main() {
  company := Company{"Google"}
  post1 := Post{"Hello World"}
  post2 := Post{"Hola Mundo"}
  user := User{"Julio", company, []Post{post1, post2}}
  fmt.Println(user.name)
  fmt.Println(user.company.name)
  fmt.Println(user.posts)
}