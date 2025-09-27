package main
import("fmt")

type Student struct{
	id int
	name string
	age int
}

func main(){
	rahim :=Student{101,"sujon",19}
	fmt.Println(rahim)
}