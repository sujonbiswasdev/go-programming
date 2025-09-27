package main
import("fmt")

func myMessage() {
  fmt.Println("I just got executed!")
}
func fullName(name string){
	fmt.Println("my name is",name)
}
func squre(num1 int){
	result:=num1*num1
	fmt.Printf("%v\n",result)
}
func main() {
	fmt.Println("funtion")
	myMessage()
	fullName("sujon biswas")
	fullName("rajon biswas")
	squre(4)
}