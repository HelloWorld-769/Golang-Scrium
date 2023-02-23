package main

import (
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
)

func check(value interface{}) error {
	str := value.(string)
	if str != "abc" {
		return errors.New("String must be 'abc'")
	}
	return nil

}
func stringsEqual(s string) validation.RuleFunc {
	return func(value interface{}) error {
		str := value.(string)
		if s != str {
			return errors.New("Strings are not equal.")
		}
		return nil
	}
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,

		//validation for Zip field
		validation.Field(&a.Street,
			validation.Required,
			validation.Length(1, 15),
		),

		//validation for Zip field
		validation.Field(&a.City,
			validation.Required,
			validation.Length(1, 10),
		),

		//validation for Zip field
		validation.Field(&a.State,
			validation.Required,
			validation.Length(1, 10),
		),

		//validation for Zip field
		validation.Field(&a.Zip,
			validation.Required,
			validation.Length(1, 10),
		),
	)

}

/*

func main() {

	//validating a  simple value...
	// data := "www.example.com"

	// err := validation.Validate(data,

	// 	validation.Required,
	// 	validation.Length(5, 100),

	// 	is.Email,
	// )
	// fmt.Println(err)

	// err := validation.Validate("sajdja", validation.By(check))
	// fmt.Println(err)

	// err = validation.Validate("xyz", validation.By(stringsEqual("abc")))

	// fmt.Println(err)

	//Validating a struct..
	a := Address{
		Street: "123 Main Stasdasd",
		City:   "Chicago",
		State:  "IL",
		Zip:    "",
	}
	err := a.Validate()
	fmt.Println(err)

}*/

type Customer struct {
	Name   string
	Emails string
}

var nameRule = []validation.Rule{
	validation.Required,
	validation.Length(1, 10),
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, nameRule...),
		validation.Field(&c.Emails, nameRule...),
	)
}
func main() {
	c := Customer{
		Name:   "",
		Emails: "",
	}
	err := c.Validate()
	fmt.Println(err)
}
