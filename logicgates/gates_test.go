package logicgates

import (
	"fmt"
	"testing"
)

var aTest = []bool{false, true}

func TestOr(t *testing.T) {
	for i := range aTest {
		for j := range aTest {
			fmt.Println("OR gate, Case", aTest[i], aTest[j], Or(aTest[i], aTest[j]))
		}
	}
}

func TestXor(t *testing.T) {
	for i := range aTest {
		for j := range aTest {
			fmt.Println("XOR gate, Case", aTest[i], aTest[j], Xor(aTest[i], aTest[j]))
		}
	}
}

func TestAnd3(t *testing.T) {
	for i := range aTest {
		for j := range aTest {
			for k := range aTest {
				fmt.Println("And3 gate, Case", aTest[i], aTest[j], aTest[k], And3(aTest[i], aTest[j], aTest[k]))
			}
		}
	}
}

func TestMultiplexor(t *testing.T) {
	for i := range aTest {
		for j := range aTest {
			for k := range aTest {
				fmt.Println("Multiplexor: a:", aTest[j], "b:", aTest[k], "sel:", aTest[i], Multiplexor(aTest[j], aTest[k], aTest[i]))
			}
		}
	}
}
func TestDeMultiplexor(t *testing.T) {
	for i := range aTest {
		for j := range aTest {
			a, b := DeMultiplexor(aTest[i], aTest[j])
			fmt.Println("TestDeMultiplexor gate, Case", aTest[i], aTest[j], a, b)
		}
	}
}
