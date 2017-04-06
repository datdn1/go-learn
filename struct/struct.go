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

type Person struct {
	Name string
}

type Employee struct {
	*Person
	Salary float32
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

	sa2 := NewSaiyan("datdn2", 1000, "Lac Long Quan, Tay Ho")
	fmt.Println("Name = ", sa2.name, "Age = ", sa2.age, "Address = ", sa2.address)

	per := &Person{Name:"datdn1"}
	fmt.Println(per.getName())

	em := &Employee{Person:per, Salary:1000.0}
	fmt.Println("Name = ", em.getName(), "Salary = ", em.Salary)

	arr1 := [10]int {1,2,3,4,5,6,7,8,9,0}
	fmt.Println(arr1)
	for index, value := range arr1 {
		fmt.Println("Index = ", index, " Value = ", value)
	}
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

/*
Factory tao instance thay cho ham khoi tao
 */
func NewSaiyan(name string, age int, address string) *Saiyan  {
	return &Saiyan{name:name, age:100, address:address}

	// or
	// return  new(Saiyan)
}

// ham cho struct person
func (p *Person) getName() string {
	return p.Name
}






