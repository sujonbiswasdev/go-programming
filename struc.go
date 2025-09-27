package main
import("fmt")

type Student struct{
	id int
	adress string
	age int
}

func displayInfo(s Student){
	fmt.Println(s.id)
	fmt.Println(s.adress)
	fmt.Println(s.age)
}

func (x *Student) increseAge(val int){
	x.age = x.age+val
}

func main(){
	rahim :=Student{101,"sylhet",19}
	rajon :=Student{100,"sylhet,bangladesh",17}
	rahim.increseAge(2)
	displayInfo(rahim)
	displayInfo(rajon)
}