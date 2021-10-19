package main

import (
	"fmt"
	"os"
)
func main() {
	
	makeItBlue()
	vTov()
	assignWithExpressions()
	findRecPer()
	multiAssign()
	multiAssign2()
	multiShortFunc()
	swapper()
	swapper2()
	fix1()
	fix2()
	fix3()
	fix4()
	fix5()
	countArgs()
	printThePath()
	printYourName()
	printMoreName()
}



func makeItBlue() {
	color:= "green"
	fmt.Println(color)
	color = "blue"
	fmt.Println(color)
}

func vTov (){
	fmt.Println("******** variable to variable ***********")
	color := "green" 

	fmt.Println(color)

	color = "dark " + color
	fmt.Println(color)

}

func assignWithExpressions() {
	fmt.Println("******** assign with expressions ***********")
	fmt.Println(3.14 * 2)
}

func findRecPer(){
	fmt.Println("******** find rectangle perimeter ***********")
	w, h := 5, 6
	
	p := 2*(w + h)

	fmt.Println(p)
}

func multiAssign () {
	fmt.Println("******** multi assign ***********")
	var (
		lang string
		version int
	)
	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}

func multiAssign2 () {
	fmt.Println("******** multi assign2 ***********")
	var (
		planet string
		isTrue bool
		temp float64
	)
	
	planet = "good"
	isTrue = true
	temp = 19.5
planet, isTrue, temp = "good", true, 19.5

	fmt.Println("air is "+planet+" on mars")
	fmt.Println( "It's",isTrue)
	fmt.Println("it is", temp, "degrees")
}
func multiShortFunc () {
	fmt.Println("******** multi short func ***********")

	_, b := multi()
	fmt.Println(b)
}

func multi()(int, int) {
	return 5, 4
}

func swapper (){
	fmt.Println("******** swapper ***********")

	color, color2 := "red", "blue"
	fmt.Println(color, color2)

	color, color2 = "orange", "green"
fmt.Println(color, color2)

}
func swapper2 (){
	fmt.Println("******** swapper 2 ***********")

	red, blue := "red", "blue"
	fmt.Println(red, blue)

	red, blue = "blue", "red"
fmt.Println(red, blue)

}
func fix1(){
	fmt.Println("******** fix 1 ***********")

	a, b := 10, 5.5
	fmt.Println(float64(a)+ b)
}
func fix2(){
	fmt.Println("******** fix 2 ***********")

	a, b := 10, 5.5
	a = int(b)

	fmt.Println(float64(a)+ b)
}
func fix3(){
	fmt.Println("******** fix 3 ***********")
a := 5.5
	
	fmt.Println(float64(int(a))+ 0.5)
}
func fix4(){
	fmt.Println("******** fix 4 ***********")
age := 2


	fmt.Println(7.5+float64(age))
}
func fix5(){
	fmt.Println("******** fix 5 ***********")
min := int8(127)
max := int16(1000)
fmt.Println(max + int16(min))

}
func countArgs(){
	fmt.Println("******** count the args ***********")
	count := len(os.Args) -1
	
	fmt.Println(os.Args)
	fmt.Printf("There are %d names.\n", count)
	
}


func printThePath()  {
	fmt.Println("******** printThePath ***********")
	fmt.Println(os.Args[0])	
}

func printYourName() {

	fmt.Println("hi "+os.Args[1])
	fmt.Println("how are you?")
}

func printMoreName(){
	fmt.Println("there are ", len(os.Args)-1,"people")
	fmt.Println("Hello great ",os.Args[1])
	fmt.Println("Hello great ",os.Args[2])
	fmt.Println("Hello great ",os.Args[3])
	fmt.Println("Nice to meet you all")
}