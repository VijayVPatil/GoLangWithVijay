package main

import "fmt"

const (
	a = iota //0
	b        //1
	c        //2
	d        //3
)

const (
	X = iota //0
	Y        //1
	Z        //2
)

const (
	M = iota //0
	N        //1
	O = 100  //100
	P        //inherits value of previous constant
	Q = iota //2 ((continues from the last iota value before manual setting))
	R        //3
	S        //4
)

//For bitwise operation
const (
	v1 = 0
	v2 = 1 << iota // 1
	v3             // 2
	v4             // 4
	v5             // 8
)

func main() {
	fmt.Printf("value a:%v and type: %T\n", a, a)
	fmt.Printf("value b:%v and type: %T\n", b, b)
	fmt.Printf("value c:%v and type: %T\n", c, c)
	fmt.Printf("value d:%v and type: %T\n", d, d)

	fmt.Printf("value X:%v and type: %T\n", X, X)
	fmt.Printf("value Y:%v and type: %T\n", Y, Y)
	fmt.Printf("value Z:%v and type: %T\n", Z, Z)
	fmt.Printf("value M:%v and type: %T\n", M, M)

	fmt.Printf("value N:%v and type: %T\n", N, N)
	fmt.Printf("value O:%v and type: %T\n", O, O)
	fmt.Printf("value P:%v and type: %T\n", P, P)
	fmt.Printf("value Q:%v and type: %T\n", Q, Q)
	fmt.Printf("value R:%v and type: %T\n", R, R)
	fmt.Printf("value S:%v and type: %T\n", S, S)

	fmt.Printf("value v1:%v and type: %T and %b\n", v1, v1, v1)
	fmt.Printf("value v2:%v and type: %T and %b\n", v2, v2, v2)
	fmt.Printf("value v3:%v and type: %T and %b\n", v3, v3, v3)
	fmt.Printf("value v4:%v and type: %T and %b\n", v4, v4, v4)
	fmt.Printf("value v5:%v and type: %T and %b\n", v5, v5, v5)

}
