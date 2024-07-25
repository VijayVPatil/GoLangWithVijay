# Generics (Type Parameters)

1. Generics help us to write flexible, reusable code such that we can use functions , types , data structures with any type.
2. Generics enable us to use DRY principle(Don't Repeat Yourself)

Let's see how this works

```
func addFloat(a,b float64)float64{
    return a + b
}

func addInt(a,b int)int{
    return a + b
}

func main(){
    fmt.Println(addInt(2,3))
    fmt.Println(addFloat(2.2,3.3))
}

```

1. In above example we see two int's and float values are added using the function addInt and addFloat.
2. This code first is repeated and can be more optimised using type parameters
3. Lets create a generic function

```
func function_name[T any](value T) {
    fmt.Println(value)
}
```

Above we had added a additional square bracket where we specify the type our generic function want to take

#### We can distribute generics in type

### Generic Function

Add below generic function in your code

```
func addTypes[T any](a, b T) T {
	return a + b
}

```

### Generics with type constraint.

Add below generic function with type constraint in your code

```
func addTypes[T int | float64](a, b T) T {
	return a + b
}

```

1. Above generic function will take either int or float hence it is type constraint
2. Please refer Day 16(3)-Generics\generics_basic_example
