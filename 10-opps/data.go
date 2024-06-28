package model

import "fmt"

var (
	CompanyName = "test"
	companyLocation = "Mangalore"
)

type Person struct {
	Name string
	age int
}

//Get age of the person
func (p *Person) GetAge() int{
 	return p.age
}

func (p *Person) getName() string {
	return p.Name
}

type company struct {

}

//GetPerson get the person object

func GetPerson() *Person {
	p := &Person{
		Name: "Swarn",
		age: 30,
	}
	fmt.Println("Model package")
	fmt.Println(p.Name)
	fmt.Println(p.age)
	return p
}

func getCompanyName() string {
    return CompanyName
}