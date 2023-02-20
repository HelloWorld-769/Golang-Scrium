package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println(today)
	fmt.Println("Format Date:", today.Format("January 02,2006"))
	fmt.Println("Format time:", today.Format("3:04 PM"))

	//loaction
	fmt.Println("Location", today.Location())

	//creating new date
	date := time.Date(2020, 01, 20, 14, 55, 48, 324359102, time.Local)
	fmt.Println("Date is :", date) //op-> Date is : 2020-01-20 14:55:48.324359102 +0530 IST

	next_date := date.AddDate(1, 2, 1) //1 year,2 months,3 days
	fmt.Println("Next date is:", next_date)

	fmt.Println("Time at present with time zone is:", today.Format(time.RubyDate))

}
