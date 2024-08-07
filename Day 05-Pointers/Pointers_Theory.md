# Pointers

We use pointers in golang to store the memory address of a type.
Pointers are stored in hexadecimal format

### Pointer declaration

Let's see how we can declare a pointer

```
package main

import "fmt"

var memOne *int

//Following will print <nil> as zero value of pointer type is <nil>
fmt.Println("The memOne is ",memOne)
```

Here we are declaring integer pointer type with name argument 'memOne'. This will store address. 
The zero value of a pointer type variabe is 'nil'  

#### Extract address of a type

Let's say we have a integer type of a variable num1. If we need to see at what memory location is the num1 stored in we can do that using '&' like below:

```
package main

import "fmt"

var num1 int
var num1Address *int
num1Address=&num1

fmt.Println("Address of num1 is ",&num1)
fmt.Println("Address of num1 is ",num1Address)
```

#### Declaring Pointers using short hand declaration

We can use short hand operator(:=) to declare a variable 'dig1' of pointer type . We do not need to mention the pointer type while declaring the variable . The type of pointer variable is inferred by the compiler .For sanity check let see what will be the type of the variable 'dig1'

```
package main

import "fmt"

var a int=23
dig1:=&a

fmt.Printf("The address of a stored in dig1 is %v and type is %v\n",dig1,dig1)
```

###### Wait! But did you ran this code on your machine. Do it and check for yourself! Come back and continue

#### Dereferecing a pointer

We can dereference a pointer that is extract the value at that address using '\*' operator.

```

package main

import "fmt"

nos := 34
nos1 := &nos

	//Using '*' we can dereference the pointer to get the value at that address
	value := *nos1
	fmt.Println("The Address is :", nos1)
	fmt.Println("The Value is at address stored in nos1 extracted using value which is:", value)
```

#### Pointer to a Pointer

We can also have a pointer to a memory location stored in a variable.
Let's breakdown the code below.

We have a variable 'a' where alue 32 is stored. We store the address of 'a' in variable 'b'. Now the address of 'b' that is of inter pointer type is stored in 'c' . Checkout the formatted print statement to check how it is dereferenced. 

```

package main

import "fmt"

func main() {
	a := 32

	var b *int = &a

	var c **int = &b

	fmt.Println("Value of a is ", a, " and value of pointer to a is that is stored in b", b)

	fmt.Printf("To access the value at address stored in b using dereferencing , the value will be %v\n", *b)

	fmt.Println("*c will give me address of pointer of a that is &a that is b", *c)
	fmt.Println("**c will give ", **c)

}

```
