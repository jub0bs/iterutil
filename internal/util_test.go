package internal_test

func alwaysFalse[E any](_ E) bool {
	return false
}

func equal[E comparable](target E) func(E) bool {
	return func(e E) bool {
		return e == target
	}
}
