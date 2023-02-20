// package main

// import "fmt"

// func main() {
// 	// log.SetPrefix("greetings: ")
// 	// log.SetFlags(0)
// 	// var err error
// 	// var message string
// 	// message, err = greetings.Hello("")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(message)

// 	//intput
// 	// fmt.Print("Enter your first name:")
// 	// var first string
// 	// fmt.Scanln(&first)
// 	// fmt.Println(first)
// 	// fmt.Print("your First name is: ", first)

// 	//Array in go
// 	// var arr [3]int
// 	// arr[0] = 1
// 	// arr[1] = 2
// 	// arr[2] = 3
// 	// var i int
// 	// for i = 0; i < 3; i++ {
// 	// 	fmt.Println(arr[i])
// 	// }

// 	//dynamic array

// 	// arr := [...]int{1, 2, 3, 4, 5, 6}
// 	// for i := 0; i < len(arr); i++ {
// 	// 	fmt.Printf("%d ", arr[i])
// 	// }

// 	// var num = 2

// 	// fmt.Printf("%T", num)

// 	// fmt.Println("Enter a day: ")
// 	// var day int
// 	// fmt.Scanln(&day)
// 	// switch day {
// 	// case 1:
// 	// 	fmt.Println("Monday")
// 	// case 2:
// 	// 	fmt.Println("Tuesday")
// 	// case 3:
// 	// 	fmt.Println("Wednesday")
// 	// case 4:
// 	// 	fmt.Println("Thursday")
// 	// case 5:
// 	// 	fmt.Println("Friday")
// 	// case 6:
// 	// 	fmt.Println("Saturday")
// 	// case 7:
// 	// 	fmt.Println("Sunday")
// 	// default:
// 	// 	fmt.Println("Invalid")
// 	// }

// 	// // create two slices
// 	// evenNumbers := []int{2, 4}
// 	// oddNumbers := []int{1, 3}
// 	// //oddNumber := 5
// 	// // add elements of oddNumbers to evenNumbers
// 	// evenNumbers = append(evenNumbers, oddNumbers...)
// 	// fmt.Println("Numbers:", evenNumbers)
// 	// fmt.Println(rand.Perm(10)).
// 	// x := 10
// 	// var y float32
// 	// fmt.Println(y)
// 	// fmt.Printf("%T", float32(x))

// }

// Golang program to illustrate
// the pointer to struct
// Golang program illustrates how
// to implement an interface
// Go program to illustrate the function
// as a field in Go structure
package main

import "fmt"

// Finalsalary of function type
type Finalsalary func(int, int) int

// Creating structure
type Author struct {
	name      string
	language  string
	Marticles int
	Pay       int
	t         Finalsalary
	// Function as a field
	salary Finalsalary
}

// Main method
func main() {

	// Initializing the fields
	// of the structure
	result := Author{
		name:      "Sonia",
		language:  "Java",
		Marticles: 120,
		Pay:       500,
		salary: func(Ma int, pay int) int {
			return Ma * pay
		},
	}

	// Display values
	fmt.Println("Author's Name: ", result.name)
	fmt.Println("Language: ", result.language)
	fmt.Println("Total number of articles published in May: ", result.Marticles)
	fmt.Println("Per article pay: ", result.Pay)
	fmt.Printf("Total salary:%T ", result.salary(result.Marticles, result.Pay))
}
