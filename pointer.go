package main
import("fmt")

func changeValue(x int){
	x = 25
}

func callByReferance(ptr *int){
	*ptr = 80
}
func main() {
	// var p *int
	// i := 42
	// p = &i
	// fmt.Println(p) 
	// fmt.Println(*p) 
	// // x:=10
	// fmt.Println(&x)

	// x:=20
	// changeValue(x)
	// fmt.Println(x)
	// callByReferance(&x)
	// fmt.Println(x)


}