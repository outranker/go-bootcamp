package main

import (
	"fmt"
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