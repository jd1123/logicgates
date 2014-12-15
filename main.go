package main

import (
	"fmt"

	"github.com/jd1123/logicgates/logicgates"
)

func main() {
	a := true
	b := false
	fmt.Println("Nand a,b:", logicgates.Nand(a, b))
	fmt.Println("Not a:", logicgates.Not(a))
	fmt.Println("a And b:", logicgates.And(a, b))
	fmt.Println("a Or b:", logicgates.Or(a, b))
}
