package main
import "fmt"
func show(name string,age int){
  fmt.Println("You must be",name,"of age",age)
}
func add(){
  var name string
  var age int
  fmt.Print("Enter your name:")
  fmt.Scanln(&name)
  fmt.Print("Enter your age:")
  fmt.Scanln(&age)
  show(name,age)
}
func main() {
  add()
}

