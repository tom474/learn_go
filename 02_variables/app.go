package main

import (
	"bufio"
	"fmt"
	"os"
)

const COUNTRY string = "Viet Nam"

var city string = "Ho Chi Minh City"
var (
	food  string = "Beefsteak"
	sport string = "Football"
)

func main() {
	var name string = "Cuong Tran"
	var age int = 20
	fmt.Println("My name is", name, "and I am", age, "years old.")
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)

	phoneNumber := "0912345678"
	fmt.Println("My phone number is " + phoneNumber + ".")

	fmt.Println("Global country:", COUNTRY)
	fmt.Println("Global city:", city)
	fmt.Println("Global food:", food)
	fmt.Println("Global sport:", sport)

	var message, firstname, lastname string
	fmt.Print("Please enter your message: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}
	fmt.Println("Your message:", message)
	fmt.Print("Please enter your firstname: ")
	fmt.Scan(&firstname)
	fmt.Print("Please enter your lastname: ")
	fmt.Scan(&lastname)
	fmt.Println("Hello " + firstname + " " + lastname + "!")

	message = fmt.Sprint("My name is ", name)
	fmt.Println(message)
	message = fmt.Sprintln("I am", age, "years old.")
	fmt.Println(message)

	name = "Cuong Tran"
	age = 20
	height := 1.78
	isGraduated := false
	progress := 90
	fmt.Printf("Data type of name: %T \n", name)
	fmt.Printf("Data type of age: %T \n", age)
	fmt.Printf("Data type of height: %T \n", height)
	fmt.Printf("Data type of isGraduated: %T \n", isGraduated)
	fmt.Printf("Data type of progress: %T \n", progress)

	fmt.Printf("My height: %.2fm \n", height)
	fmt.Printf("My height: %.5fm \n", height)
	fmt.Printf("My height: %.1fm \n", height)
	fmt.Printf("Graduated: %t \n", isGraduated)
	fmt.Printf("Progress: %d%% \n", progress)
}
