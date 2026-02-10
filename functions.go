package main

//func FunctionName(param1 type, param2 type, param3 type) (result1, result2 int) {
//   code to be executed
//}

// FUNCTION AS A VALUE
//could be useful for not duplicating a lot of functions with similar code parts, for example we could
//create function calculate, pass inside our values and exact function to calculate without need of spamming
// if-else statements inside global "calculate" function like below

//func compute(operation func(float64, float64) float64) float64 {
//	return operation(3, 4)
//}
//sum := func (a,b) int { return a+b }
//result := compute(sum(a,b))
//instead of if "sum".. if "multiply"..... we can pass exact operation while initiating

// CLOSURES
//func intSeq() func() int {
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
//}
//nextInt := intSeq() // counter "i" created for the first time and live until be redeclared
//fmt.Println(nextInt()) // 1  changes counter "i" that lives inside the nextInt like a struct field
//fmt.Println(nextInt()) // 2
//fmt.Println(nextInt()) // 3
