package logicgates

// Axiomatically exists
func Nand(a, b bool) bool {
	return !(a && b)
}

// Build other gates from Nand and Nand alone
func Not(a bool) bool {
	return Nand(a, a)
}

// Not follows from being built with Nand
func And(a, b bool) bool {
	return Not(Nand(a, b))
}

func Or(a, b bool) bool {
	return Nand(Nand(a, a), Nand(b, b))
}

func Xor(a, b bool) bool {
	return Or(And(a, Not(b)), And(b, Not(a)))
}
