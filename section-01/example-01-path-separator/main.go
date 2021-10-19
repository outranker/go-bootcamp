package main

import (
	"fmt"
	"path"
)
func main() {
	separatePath("src/assests/image.jpg")
	separatePathCoolOne("src/assests/image.jpg")
}


func separatePath(input string)  {
	var dir, file string
	dir, file = path.Split(input)
	fmt.Println("dir :",dir)
	fmt.Println("file:",file)


}

func separatePathCoolOne(input string)  {
	
	_, file := path.Split(input)
	
	fmt.Println("file:",file)
}
