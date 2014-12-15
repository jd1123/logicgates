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
	return Nand(Nand(a, b), Nand(a, b))
}

func Or(a, b bool) bool {
	return Nand(Nand(a, a), Nand(b, b))
}
