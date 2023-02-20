package main

import "fmt"

//in interface we can not define variable . interface ke instace nhi bn skta...
//int struct we can define function i=but it is complex task so we define it using interfaces

type Human struct {
	Name     string
	Age      int
	Dialogue string
}

type Monster struct {
	Type  string
	Age   int
	Power string
}

type Being interface {
	tellAge()
}

func ageTeller(b Being) {
	b.tellAge()
}

func (p *Human) tellAge() {

	fmt.Printf("I'm a human and I'm %v years old.\n", p.Age)
}

func (m *Monster) tellAge() {
	fmt.Printf("I'm a Monster and I'm %v years old.\n", m.Age)
}

func main() {
	Thanos := Monster{Type: "Titan", Age: 1232, Power: "Gauntlet"}
	//	Tony := Human{Name: "Tony", Age: 42, Dialogue: "Assemble"}

	// ageTeller(&Thanos)
	// ageTeller(&Tony)
	Thanos.tellAge()
}
