package main

import (
	"fmt"
)

/*
Dinh nghia struct
 */
type Saiyan struct {
	name string
	age int
	address string
}





func main() {
	x := getPower()

	name, age := "datdn11111111", 100
	fmt.Println(x)
	fmt.Printf("Name = %s Age = %d\n", name, age)

	limit := 6
	message, nameIsValid := checkName(name, limit)
	fmt.Println(message, nameIsValid)

	sa := Saiyan{name:"datdn1", age:100, address:"291 Lac Long Quan, Cau Giay, HN"}
	fmt.Println(sa)

	fmt.Println("Change age of saiyan");
	super(&sa)
	fmt.Println("Age = ", sa.age)

	sa1 := &Saiyan{name:"khoidn1", age:80, address:"Lac Long Quan"}
	sa1.supper1()
	fmt.Println("SA1 = ", sa1.age)
}

/*
Func khong co tham so truyen vao
 */
func getPower() int {
	return 9000
}

func checkName(name string, limit int) (string, bool){
	var message string
	var isValid bool

	if len(name) > limit {
		message = "InValid"
	} else {
		message = "Valid"
	}
	isValid = len(name) <= limit
	return message, isValid
}

func super(s *Saiyan)  {
	s.age += 100
}

func (s *Saiyan) supper1() {
	s.age += 100
}



