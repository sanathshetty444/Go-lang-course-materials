package main

import "fmt"

type contactInfo struct {
	email string
	zip   string
}
type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	p := person{
		firstname: "Sanath",
		lastname:  "Shetty",
		contact: contactInfo{
			email: "sa",
			zip:   "ss",
		},
	}
	p.print()
	demo := &p
	demo.updateFirstName("Demo")
	p.print()
}

func (p *person) updateFirstName(firstname string) {
	(*p).firstname = firstname
}
func (p person) print() {
	fmt.Printf("%+v \n", p)
}
