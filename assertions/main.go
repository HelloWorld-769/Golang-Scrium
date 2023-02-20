package main

import "fmt"

type ninjaStar struct {
	owner string
}
type ninjaSword struct {
	owner string
}
type ninjaWeapon interface {
	attack()
}

func (n ninjaStar) pickup() {
	fmt.Println("Picking up the ninja star..")
}
func (n ninjaStar) attack() {
	fmt.Println("Throwing ninja star..")
}
func (n ninjaSword) attack() {
	fmt.Println("Throwing ninja sword..")
}

func main() {

	//doubt hai
	var nin ninjaWeapon

	nin = ninjaSword{"HelloWorld"}
	fmt.Println(nin.(ninjaSword).owner)

	weapons :=
		[]ninjaWeapon{
			ninjaStar{"HelloWorld"},
		}

	for _, weapon := range weapons {
		weapon.(ninjaStar).attack()

	}
}
