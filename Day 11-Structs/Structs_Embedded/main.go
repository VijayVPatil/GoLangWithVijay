package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	audi := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Steel",
		model: "X1 Genz",
		doors: 2,
		color: "Green",
	}

	maruti := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Aluminium",
		model: "Suzuki 800",
		doors: 4,
		color: "Red",
	}

	fmt.Printf("%#v\n", audi)
	fmt.Printf("%#v\n", maruti)

	fmt.Println(audi.color)
	fmt.Println(audi.model)
	fmt.Println(audi.doors)
	fmt.Println(audi.make)
	fmt.Println(audi.engine)
	fmt.Println(audi.engine.electric)

	fmt.Println(maruti.color)
	fmt.Println(maruti.model)
	fmt.Println(maruti.doors)
	fmt.Println(maruti.make)
	fmt.Println(maruti.engine)
	fmt.Println(maruti.engine.electric)
}
