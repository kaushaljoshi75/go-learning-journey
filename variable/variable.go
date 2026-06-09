package variable

import "fmt"

func  DeclareVariable() {

	// Variable declaration and initialization in Go for integer type
	intVariables()

	// Variable declaration and initialization in Go for string type
	stringVariables()

	// Variable declaration and initialization in Go for boolean type
	booleanVariables()

	// Variable declaration and initialization in Go for unsigned integer type
	uintVariables()

	// Variable declaration and initialization in Go for float type
	floatVariables()

	// Variable declaration and initialization in Go for complex type
	complexVariables()

	// Variable declaration and initialization in Go for byte type
	byteVariables()

	// Variable declaration and initialization in Go for rune type
	runeVariables()

	// Constants in Go
	constants()
}

// In Go, you can declare variables using the var keyword, and you can also use short variable declaration with := for type inference. Here are examples of both methods for nume~ric, boolean, and string types.


// Signed Integer variables
func intVariables() {
	var age int
	var year int = 2025
	currentYear := 2026

	fmt.Println("Age:", age)
	fmt.Println("Year:", year)
	fmt.Println("Current Year:", currentYear)
}

// String variables
func stringVariables() {
	var name string
	var lang string = "Go Programming"
	firstName := "John"

	fmt.Println("Name:", name)
	fmt.Println("Language:", lang)
	fmt.Println("First Name:", firstName)
}

// Boolean variables
func booleanVariables() {
	var isGoFun bool
	var isProgrammingFun bool = true
	isLearningFun := true

	var a = true
	var b = false

	c := a && b // Logical AND operation
	d := a || b // Logical OR operation

	fmt.Println("Is Go fun?", isGoFun)
	fmt.Println("Is programming fun?", isProgrammingFun)
	fmt.Println("Is learning fun?", isLearningFun)
	fmt.Println("Logical AND:", c)
	fmt.Println("Logical OR:", d)
}


//  Unsigned integer variables
func uintVariables() {
	var positiveNumber uint
	var anotherPositiveNumber uint = 100
	positiveShort := uint(200)

	fmt.Println("Positive Number:", positiveNumber)
	fmt.Println("Another Positive Number:", anotherPositiveNumber)
	fmt.Println("Positive Short:", positiveShort)
}


// Float variables
func floatVariables() {
	var pi float64
	var e float32 = 2.7182834342343434343324234343
	piShort := 3.14159

	fmt.Println("Pi:", pi)
	fmt.Println("Euler's Number:", e)
	fmt.Println("Pi Short:", piShort)
}

// Complex variables
func complexVariables() {
	var complexNumber complex128
	var anotherComplexNumber complex64 = 1 + 2i
	complexShort := 3 + 4i

	fmt.Println("Complex Number:", complexNumber)
	fmt.Println("Another Complex Number:", anotherComplexNumber)
	fmt.Println("Complex Short:", complexShort)
	fmt.Printf("Real part of complexShort: %f\n", real(complexShort))
	fmt.Printf("Imaginary part of complexShort: %f\n", imag(complexShort))
}

// Bytes variables
func byteVariables() {
	s := "Hello"
	b := []byte(s)          // convert string → bytes
	fmt.Println(b) 
}


// Rune variables
func runeVariables() {

	var hindi rune = 'न'
    var heart rune = '❤'

	fmt.Printf("Hindi Rune: %c, Unicode code point: %U\n", hindi, hindi)
	fmt.Printf("Heart Rune: %c, Unicode code point: %U\n", heart, heart)

	s := "héllo"

    // range loop → gives you RUNES (real characters) ✅
    for i, r := range s {
        fmt.Printf("%d: %c\n", i, r)
    }

	r := 'A' // A rune literal
	fmt.Printf("Rune: %c, Unicode code point: %U\n", r, r)
}

// Constants in Go are immutable values that cannot be changed after they are defined. They are declared using the const keyword and can be of any basic data type, including numeric, string, and boolean types. Constants can also be untyped, which allows them to be used in a wider range of contexts without requiring explicit type conversion.
func constants() {
	const Pi = 3.14159
	const Greeting = "Hello, World!"
	const IsGoFun = true

	type myBool bool
	const IsLearningFun myBool = true

	fmt.Println("Pi:", Pi)
	fmt.Println("Greeting:", Greeting)
	fmt.Println("Is Go fun?", IsGoFun)
	fmt.Println("Is learning fun?", IsLearningFun)	
}