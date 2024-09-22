### Q1. What are packages in golang and why do we use them?

1. Packages in golang is the way to organise and reuse the code.
2. It is a collection of files.
3. Packages help to modularize code, making it easier to manage, understand, and maintain.
4. Each package provides specific functionalities, which can be imported and used in other packages or applications.

There are types of packages in golang:

1. Core packages : Core packages are provided by go code or standard library. Ex: fmt , os
2. Third-Party packages : Packages that we import from github. Ex: Gorm , Mysql
3. User defined Packaged : If we are follwoing some pattern like MVC pattern of three layer pattern , we will be distribute in small packages. More over used to maintain the code

### Q2. What are variables and constant in Golang?

###### Constants

1. Constants in Go are immutable values that are known at compile-time and cannot be changed during the execution of the program.
   a. we can declare constant using 'const' key word.
   b. Constants can be declared both inside and outside of a function
   c. In Go, constants cannot be initialized using the short-hand declaration operator (:=).

- Properties of Constants :

  - a. Compile-time Determination: Constants are evaluated at compile-time.
  - b. Immutability: Once set, their value cannot be changed.
  - c. Types of constant:
    - i. Typed Constants:
      Constants declared with Type  
      const ABC int=2
    - ii. Untyped Constants:
      const DEF="Navy"
  - d. Used in enumeration (IOTA):
    'IOTA' is a predeclared identifier used in the constant declaration to simplify the creation of incrementing values.It resets to 0 whenever the const keyword appears and increments by 1 for each subsequent constant in the block.
    Check Example code : zz-Interview-Prep\01_Basic\Q2_iota\main.go

```
const(
    a =iota
    b
    c
    d
)
```

- e. Bitwise Operations
  Bit Flags: Constants are often used in bitwise operations to define flags.

```
const (
    v1  = 0
    v2  = 1 << iota // 1
    v3             // 2
    v4             // 4
    v5             // 8
)
```

2. Variables in golang , we can change the value through out the code.

### Q3. Can we return multiple values from the function in go?

Yes we can return multiple values in golang. There are two ways to return multiple values in golang. We can send comma seprated values with return statement

##### Basic Syntax:

```
func functionName(params)(returnType1,returnType2,...){
  //function body
  return value1,value2,....
}

```

Lets take a example of Swap Two integers

```
func main(){
  a,b:= 3,4
  swapInts(a,b)
}

func swapInts(a ,b int)(int , int){


     return b,a
}

```

##### Named return values

```
func main(){

}

```

### Q4. What are string literal?

String literal are used to represent sequence of characters .
Two Types of String Literals,

- Interpreted String Literal:
  1.  Interpreted String literals are enclosed in double quotes(`"`)
  2.  Within these the new line , tabs , Unicode characters are represented as escape sequences.

```
   interpolated := "Hello, World!\n"
   unicode := "Gophers are \u263A"
```

Common escape sequences:

\n for newline
\t for tab
\" for double quote
\\ for backslash
\uXXXX for Unicode characters
