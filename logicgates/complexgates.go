package logicgates

func And3(a, b, c bool) bool {
	return And(a, And(b, c))
}

func Or3(a, b, c bool) bool {
	return Or(a, Or(b, c))
}

func Multiplexor(a, b, sel bool) bool {
	return Or(And(Not(sel), a), And(sel, b))
}

func DeMultiplexor(in, sel bool) (bool, bool) {
	return And(Not(sel), in), And(sel, in)
}
