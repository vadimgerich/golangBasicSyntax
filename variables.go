package main

import "fmt"

var (
	a1 int
	a2 string
	a3 float32
) // can declare vars in separate block for better style

func main() {
	variableName := "word2"      // can be used ONLY INSIDE the functions
	var i1, i2, i3 int = 1, 2, 3 //multiple vars declaration
	var str1, int2 = "word", 5   //if type is not specified we can assign different types inline

	fmt.Println(str1, int2, i1, i3, i2, variableName)
}
