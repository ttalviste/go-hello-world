package main

import (
	"fmt"

	apiClient "github.com/ttalviste/go-hello-world/rest-api-client"
)

func main() {
	apiClient.FetchSandboxTokens()

	fmt.Println("Hello world!")

	fmt.Println("Let's work with primitive types:")
	fmt.Println("Here's an int:")
	var i int
	i = 42
	fmt.Println(i)

	fmt.Println("Here's a float32:")
	var f float32 = 3.14
	fmt.Println(f)

	fmt.Println("Here's a string:")
	_string := "Toomas"
	fmt.Println(_string)

	fmt.Println("Here's a bool:")
	b := true
	fmt.Println(b)

	fmt.Println("Here's a complex:")
	c := complex(3, 4)
	fmt.Println(c)

	fmt.Println("Here's a complex type break down in real and imaginary parts:")
	r, im := real(c), imag(c)
	fmt.Println(r, im)

	fmt.Println("Let's work with pointers:")

	fmt.Println("Creating an pointer value:")
	var firstName *string = new(string)
	// *<variable> is for dereferencing
	*firstName = "Toomas"
	fmt.Println(firstName)
	fmt.Println(*firstName)

	firstName_2 := "Toomas"
	fmt.Println(firstName_2)
	fmt.Println("Getting an pointer for variable:")
	// & - address of operator
	ptr := &firstName_2
	fmt.Println(ptr, *ptr)

	firstName_2 = "Tormis"
	fmt.Println(ptr, *ptr)

	fmt.Println("Thank you for your attention, it's Go time now!")

}
