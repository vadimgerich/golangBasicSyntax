package main

func main() {
	//ARRAYS
	//var array = [length]datatype{values}
	//var array = [...]datatype{values}    // no fixed size

	//fmt.Println(array)    -  will print all the values "[1 2 3]"
	//fmt.Println(array[0]) -  access to the "0" element of an array
	//array[1] = 10 		-  change "1" element of an array

	//var array = [5]int{1,2,3} - last 2 values would be dafault for the type
	//fmt.Println(array)  -  "[1,2,3,0,0]"

	//array := [5]int{1:10,2:40} - [0,10,0,40]  - we can set specific elements values

	//len(array) - returns length of an array

	//SLICES
	//slice_name := []datatype{values}
	//myslice := []int{}
	//slice_name := make([]type, length)

	//len(slice) - returns number ofelements
	//cap(slice) - returns capability of slice

	//slice[0] - 		first element of the slice
	//clice[2] = 10 	change value of specific element
	//slice_name = append(slice_name, element1, element2, ...)  - add an element at the end of the slice
	//slice_name = append(slice1, slice2...) 					- add elements of one slice to the end of another

	//s = s[1:4]           - вибираємо елементи з індексами 1..3 (4 не включно).
	//copy(destArr,srcArr) - copy values from src to dest

	//MAPS
	//var a map[KeyType]ValueType
	//var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
	//b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	//var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}

	//var a = make(map[KeyType]ValueType)
	//b := make(map[KeyType]ValueType)

	//val, ok :=map_name[key]    - check if element exists by value
	//delete(m, key)

}
