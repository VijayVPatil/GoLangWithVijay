# Variables

Declaring a variable: Variables are declared using the var keyword. We can have variable declaration with var in function level as well as var level . For example, to declare a variable called number of type int, you would write:  
var <Vairiable_Name> <Data_Type>

For ex:  
var number int  
var temperature float64=75.4566

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
### Short Assignment operator (Type Inference):

Inside a function (even the main function), the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable based on the value.  
So we were declaring variables as

var empty string

now with short assignment operator

empty:="" …. This is how short assignment operator is used

#### Crucial Note : When to use Var Keyword and Short Hand Declaration operator?

1. If we want zero value declaration or or just declare the variable and assigned the value in future , use the keyword var
2. When variables are to be declared and initialized at same time, use the short variable declaration operator.

### Another way to declare the variable:

A dynamic type variable declaration requires the compiler to interpret the type of the variable based on the value passed to it. The compiler does not require a variable to have type statically as a necessary requirement.  
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/de7b51ff-c7c4-4894-ae7a-6d58037aebaf)

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/24c6bca0-f687-4687-a9ba-a2529b0d07e7)

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

Examples:(Please refer the code sample attached in repository
![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/cd62a4d2-c4b1-45b8-a4fe-ae5b6b694bca)

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/5af2baa4-d040-4b3c-af45-782a47695e9c)

#### String Type

A string type represents the set of string values.
Below is how we declare string

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/c701225c-8c6b-46ea-a8cb-fb4f99bebbe4)

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/3cb89443-83e2-42ff-9014-529cc3eb6e5e)

#### Boolean Type

A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is bool; it is a defined type.

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/2d0ef683-6a6d-44ee-959a-1928780c7174)

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/c389fec3-a150-4f50-b816-78942c04ffde)

### Go Constants

If a variable should have a fixed value that cannot be changed, you can use the const keyword.  
The const keyword declares the variable as "constant", which means that it is unchangeable and read-only. Which means when constant is declared it can not be changed later

Rules:-  
1)Constant names follow the same naming rules as variables
2)Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
3)Constants can be declared both inside and outside of a function

Types Constants: Constants declared with Type
Const ABC int=1

Untyped Constants: Declared without type

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/7a16a5c5-a952-4cf3-95d1-767b77b01bd7)

![image](https://github.com/VijayVPatil/GoLangWithVijay/assets/76161912/cc4f74d0-633d-4736-b01a-a2d58a1e4b5d)
