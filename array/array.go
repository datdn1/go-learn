package main

import (
	"fmt"
)

// array trong go la gia tri khong phai la con tro tro den
// phan tu dau tien
func main()  {
	// literal array
	b := [2]string {"datdn1", "khoidn1"}
	fmt.Println(b)

	b1 := [...]int {1,2,3,4}
	fmt.Println("Length = ", len(b1))

	// slice
	arr := []int {1,2,3,4}   // tao kieu thong thuong
	fmt.Println(arr)
	PrintArr(arr)

	arr1 := make([]int, 10, 12)   // length 10, cap 12
	fmt.Println("Length = ", len(arr1), " Cap = ", cap(arr1))

	var arr2 []int    // slice nil length = 0, cap = 0
	fmt.Println(len(arr2), cap(arr2))

	arr3 := []byte {'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(arr3[1:4])  // tao slice bang cach slice mang
	PrintArrByte(arr3[1:4])

	arr4 := []string {"a", "b", "c", "d"}
	fmt.Println(arr4[:3])  // ["a", "b", "c"]
	fmt.Println(arr4[2:3]) // ["c"]
	fmt.Println(arr4[:])  // all

	// https://blog.golang.org/go-slices-usage-and-internals
	arr5 := make([]string, len(arr4), cap(arr4) * 2)
	for i := range arr4 {
		arr5[i] = arr4[i]
	}
	fmt.Println("Len = ", len(arr5), "Cap = ", cap(arr5))
}

func PrintArr(arr []int) {
	for index, value := range arr {
		fmt.Println("(", index, ", ", value, ")")
	}
}

func PrintArrByte(arr []byte) {
	for index, value := range arr {
		fmt.Println("(", index, ", ", value, ")")
	}
}

//func appendInt(src []int, data ...int) []int  {
//	numberRemaind := cap(src) - len(src)
//	numberAdd := len(data)
//	if numberAdd > numberRemaind {
//		numberNeedAlloc := numberAdd - numberRemaind
//		newSlice := make([]int, cap(src) + numberNeedAlloc)
//		copy(newSlice, src)
//		src = newSlice
//	}
//}



