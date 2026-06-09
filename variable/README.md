# Day 2 : Variables, constants, basic types


## Basic Data Types

- In go we have data types like bool, numeric, string which are called basic data types
- In numeric data types we have "int8, int16, int32, int64 and int" this are singed datatype there are unsigned version alo "uint8, uint16, uint32, uint64, and uint"
- In numeric data types we have "float32" and "float64" data types also
- In numeric data types we have "complex64" and "complex128" data types also 
- In numeric data types we have byte
- In numeric data types we have rune this datatype store the unicode point value


### bool
- This data type represent the boolean value like true or false
- It can be written with "&&" and "||" operator

```go
package main

func main() {
    a := true
    b := false

    c := a && b
    d := a || b
}
```

### Signed integers
- The singed integers will have data types like "int", "int8", "int16", "int32", "int64"
- int8 will store the value of 8 bit as signed value so it can be positive and negative with the range of -128 to 127
- int16 will store the value of 16 bit same like int8 with the range of -32,768 to 32,767
- int32 will store the value of 32 bit same like int8 , int16 with the range of -214,74,83,648 to 214,74,83,647
- int64 will store the value of 64 bit same like int8, int16, int32 with the range of -922337203685,47,75,808 to 922337203685,47,75,807
- int represents 32 or 64 bit integers depending on the underlying architecture. You will always use this as integer unless there is specific sized integer need.

```go
package main

func main() {
    var a int = 18
}
```

- Here it will decide by their own based on the system which need to consider 32 bit or 64 bit when we print this using "unsafe" package which should be used with care  as the code using it might have portability issues. So if the output is 8bytes then it is 64 bits and if 4bytes then it is 32 bits.

###  Unsigned integers

- The unsinged integers will have data types like "uint", "uint8", "uint16", "uint32", "uint64" and it is used to store positive integers only.
- uint8 will store the value of 8 bit as signed value so it will be positive with the range of 0 to 255
- uint16 will store the value of 16 bit same like int8 with the range of 0 to 65535
- uint32 will store the value of 32 bit same like int8 , int16 with the range of 0 to 4294967295
- uint64 will store the value of 64 bit same like int8, int16, int32 with the range of 0 to 18446744073709551615
- uint represents 32 or 64 bit integers depending on the underlying architecture.
- Unsigned integers are used in places where negative values are not applicable.

```go
package main

func main() {
    var a uint = 18
}
```

### Floating point types

- This data type will used to store the decimal values or fractional value
- The floating will have data types like "float32", "float64" 
- float32 will round off the value and lost precision as it simply can't hold all those digits, so it rounds.
- By default float64 will be used and float32 will only used when you specifically need to save memory (e.g. huge arrays, graphics, GPU data)
- We cant perform any action of 2 variables like one variable is int and another is float then we cant perform action on two without converting any one variable.
- floats are not exact value in all the language not just Go. Decimals are stored in binary, and some values can't be represented perfectly.
- That tiny error is why you should never use floats for money Use integers (count in cents/paise) or a decimal library instead
- Comparing floats with == is risky for the same reason.

```go
package main

func main() {
   var a float32 = 3.14
   var b float64 = 3.14159265358979
}
```

### Complex types

- This are the type which store the Complex numbers that have a normal part plus an "imaginary" part, written like 3 + 4i. They come from math/engineering.
- complex64 is smaller, less precise and complex128 will have  bigger, more precise which is default also.
- For daily use this will not be used and it will only used when you are working on some scientific maths 
- real() and imag() are built-in functions that pull out each half value before "+" it is real() value and after "+" is imag() value

```go
package main

func main() {
   var a complex128 = 3 + 4i 
}
```

### byte
- byte is just another name for uint8,an integer from 0 to 255.
- byte, it means "this is raw data," not "this is a number you'll do math on.
- Go string is really a sequence of bytes and lots of things (files, network data, request bodies) come to you as bytes.
- The most common form you'll see is a byte slice, written []byte
- That string ↔ []byte conversion is everywhere in real Go code reading files, handling HTTP request/response bodies, JSON, etc. all work with []byte.

```go
package main

func main() {
  s := "Hello"
  b := []byte(s)  
}
```


### rune 
- rune is Go's type for a single Unicode character (technically a "code point")
- It's an alias for int32, because under the hood a character is stored as a number
- single quotes make a rune 'A' is a rune 
- double quotes make a string "A" is a string
- That's a real distinction in Go, unlike JS/PHP where quote style doesn't matter.
- Where rune becomes genuinely useful is non-ASCII characters like emoji, accented letters, other alphabets. These take more than one byte, and that's the whole reason rune exists.
- The most important place you'll meet rune is looping over a string.
- A Go string is really a sequence of bytes, and a character like न or ❤ takes several bytes. So there are two ways to loop, and they behave differently
- The range loop correctly walks character-by-character, even when a character spans multiple bytes. If you instead index the string directly (s[i]), you get individual bytes, which splits multi-byte characters apart and gives you garbage for anything non-ASCII.
-  You can count real characters (not bytes) by converting to a rune slice
-  byte (uint8) is for raw data, rune (int32) is for human-readable characters, and the moment you deal with anything beyond plain English text, you reach for rune.

```go
package main

import fmt

func main() {
    s := "héllo"

    // range loop → gives you RUNES (real characters) ✅
    for i, r := range s {
        fmt.Printf("%d: %c\n", i, r)
    }

    var hindi rune = 'न'
    var heart rune = '❤'
}
```

### String types

- Strings are a collection of bytes in Go
- we can assume a string to be a collection of characters

```go
package main

func main() {
  var name string = "Go Lang" 
}
```

### Type Conversion
-  Go is very strict about explicit typing
-  There is no automatic type promotion or conversion
- To perform any action we must need to convert any variable with another type both a and b must be of the same type

```go
package main

func main() {
    a := 80           //int
	b := 91.8         //float64
	sum := a + int(b) //int + float64 not allowed
}
```


## Constant

-  Constants in Go is used to denote fixed static values.
- Constants are generally used to represent values that do not change throughout the life time of an application

### Declaring a constant
- The keyword const is used to declare a constant in Go

```go
package main

func main() {
    const pi = 3.14
}
```

### Declaring a group of constants
- There is also another syntax to define a group of constants using a single statement

```go
package main

func main() {
    const (
		retryLimit = 4
		httpMethod = "GET"
	)
}
```
- Constants, as the name indicate, cannot be reassigned again to any other value
- The value of a constant should be known at compile time. Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.

```go
package main

func main() {
    var a = math.Sqrt(4) //allowed
	fmt.Println(a)
	const b = math.Sqrt(4) //not allowed
	fmt.Println(b)
}
```

### String Constants, Typed and Untyped Constants
- Any value enclosed between double quotes is a string constant in Go For example, strings like "Hello World", "Sam" are all constants in Go. This is untyped constants.
- A string constant like “Hello World” does not have any type.
- const defines a fixed value; type defines a new named type.

```go
package main

func main() {
    const trueConst = true   // declares a CONSTANT (a value)
    type myBool bool         // declares a TYPE (a new named type)
}
```


## Variables 

- In go we used the variables to store the value of a specific type.


### Declaring a single variable 

- var name <data type> 

```go
package main

func main() {
    var name sting 
}
```

- when you declare any variable without giving the value, "Go automatically initializes it with the zero value of the variable’s type". So for int it will be 0 and for string it will be "  ".

- later same like other language we can overwrite the value of the variable 

```go
package main

func main() {
    var name sting 
    name = "Go Lang"
}
```

### Declaring a variable with an initial value

- var name <data type> = initial value

```go
package main

func main() {
    var name sting = "Go Lang"
}
```

### Type Inference

- When you declare the variable it is not compulsory to define the type at the time of declaring the variable. Go will automatically be able to infer the type of that variable using the value that you provide at initial time.

- Here like other language "php, javascript....", Go does not do implicit type coercion. So if you declare any variable with type int then you cant store string value in that variable


```go
package main

func main() {
    var name = "Go Lang"
}
```

### Multiple variable declaration

- var name, age = "initial value", "initial value" (this will use the concept of type inference)


```go
package main

func main() {
    var name, age = "Go Lang", 100
}
```

- var price, quality int = "initial value", "initial value"

```go
package main

func main() {
    var price, quality = 100, 100
}
```

- Sometime there will be case we need to declare the variables with different type as multiple, that case we can do using `var ()` inside this bracket we can declare multiple variables with different data types.

```go
package main

func main() {
    var (
        name = "Go lang"
        age = 10
        height int = 2
        weight int
    )
}
```

### Short hand declaration

- Go provide one more method to declare and initialized the variable which is called "short hand" declaration and it use this symbol  ":=" operator.

- name := initial value

- here we dont have to pass the datatype go will automatically infer the datatype based on the assign value to the variable 

```go
package main

func main() {
    name := "Go Lang"
    age := 30
}
```

- You can also declare multiple variable using "short hand" method 

```go
package main

func main() {
    name,age := "Go Lang", 30
}
```

- You always need to declare and initialized the value when you use "short hand" method as if you have declare the variable but not initialized the value then it will give error

```go
package main

func main() {
    name, age := "Go" //error
}
```

- When you use short hand syntax, then at least one of the variables on the left side of := should be newly declared. If you redeclare same variable then it will give error

```go
package main

func main() {
    name, age := "Go" , 8 
    age, height := 2, 4

    name, age := "Go" , 8  // this will give error
}
```

- You can assigned the value of variables in Go on computing during run time.

```go
package main

import math

func main() {
    a, b := 10, 11
    c := math.Min(a,b)
}
```








