package pointers

// Pointer returns a pointer to the value passed in. Pretty much just a convenience function.
func Pointer[T any](value T) *T {
	return &value
}
