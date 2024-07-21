# Variables

1. Declaring a variable: Variables are declared using the var keyword.
2. We can have variable declaration with var in function level as well as package level .
3. For example, to declare a variable called number of type int, you would write:

```
var <variable_name> <data_type>
```

For ex:

```
var number int
var temperature float64=75.4566

```

Example Code :

```
package main

import "fmt"

var a int
var b string=12
func main() {
fmt.Println("Hello World!!!!")
}
```

### You know what, we can omit the type of the variable let's see how

1. We can declare the variable by omitting type if int initalizer is present and the type of the variable will be inferred based on the initalizer if not explicitly specified.
2. But what is initalizer? an initializer in Go provides the initial value for a variable at the time of its declaration.

```
//Declare a string type variable
var name="James"
```

### We can declare multiple variables at once

```
   var b, c int = 1, 2

```

### Short Assignment operator (Type Inference):

Inside a function (even the main function), the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable based on the value.  
So we were declaring variables as

```
var empty string
```

now with short assignment operator

empty:="" …. This is how short assignment operator is used

#### Crucial Note : When to use Var Keyword and Short Hand Declaration operator?

1. If we want zero value declaration or or just declare the variable and assigned the value in future , use the keyword var

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

2. When variables are to be declared and initialized at same time, use the short variable declaration operator.

### Another way to declare the variable:

A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. The compiler does not require a variable to have type statically as a necessary requirement.

```
//Declaring a variable of int type using short hand declaration

i:=1
```

### Types

#### Numeric type

An integer, floating-point, or complex type represents the set of integer, floating-point, or complex values, respectively. They are collectively called numeric types. The predeclared architecture-independent numeric types are:

uint8 the set of all unsigned 8-bit integers (0 to 255)  
uint16 the set of all unsigned 16-bit integers (0 to 65535)  
uint32 the set of all unsigned 32-bit integers (0 to 4294967295)  
uint64 the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8 the set of all signed 8-bit integers (-128 to 127)  
int16 the set of all signed 16-bit integers (-32768 to 32767)  
int32 the set of all signed 32-bit integers (-2147483648 to 2147483647)  
int64 the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32 the set of all IEEE-754 32-bit floating-point numbers  
float64 the set of all IEEE-754 64-bit floating-point numbers

complex64 the set of all complex numbers with float32 real and imaginary parts  
complex128 the set of all complex numbers with float64 real and imaginary parts

byte alias for uint8  
rune alias for int32

Examples: Please refer the code sample attached in repository

##### Integer Type

##### String Type

A string type represents the set of string values.
Below is how we declare string

```
//Using var keyword
var name string="Jane"

//Using short hand declaration
address:="Street 5 , NYC"

```

#### Boolean Type

A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is bool; it is a defined type.

```
//Using short hand declaration
isBoll:=true

//Using Var keyword
var isThisBoolean bool = true
```

#### Rune type

### Go Constants

If a variable should have a fixed value that cannot be changed, you can use the const keyword.  
The const keyword declares the variable as "constant", which means that it is unchangeable and read-only. Which means when constant is declared it can not be changed later

Rules:-

1. Constant names follow the same naming rules as variables
2. Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
3. Constants can be declared both inside and outside of a function

Types Constants: Constants declared with Type

```
const ABC int=1
```

Untyped Constants: Declared without type

```
const DEF="Navy"
```
