package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
  fmt.Println("Hello sujon!")
  fmt.Println(30)
  // sujon biswas


  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
  // input : Scan,Scanln,Scanf
  //output : Print,Println,Printf

  var num1,num2 int
  fmt.Print("enter two number : ")
  fmt.Scan(&num1,&num2)
  fmt.Printf("num1 = %v, num2=%v\n",num1,num2)


  var fullname string
  fmt.Print("enter your name")
  fmt.Scan(&fullname)
  fmt.Printf("%v is a student\n",fullname)

}