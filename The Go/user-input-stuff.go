package main

import "fmt"
import "time"
var name string = "Greg"
var name2 string = "greg"
var username string
func main() {
  fmt.Println("Hello! My name is " + name +"! What's yours?")
fmt.Scan(&username)
if (username == name || username == name2) {
  fmt.Println("My name is", name, "Too!")
} else {
  fmt.Println("Nice to meet you,", username + "!")
}
fmt.Println("I'm currently learning C++, Python and GO!")
fmt.Println("Idk why I'm learning go tbh :sob:")
fmt.Println("Anyway, take your current time sob")
fmt.Println(time.Now())
}
