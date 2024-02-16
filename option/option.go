package option

// Option is a generic type that can be T or None
type Option[T any] interface {
	IsNone() bool
	IsSome() bool
}

// None is a type that represents the absence of a value
type none struct{}

// IsNone returns true if the Option is None
func (n none) IsNone() bool {
	return true
}

// IsSome returns false if the Option is None
func (n none) IsSome() bool {
	return false
}

// some is a type that represents the presence of a value
type some[T any] struct {
	Value T
}

// IsNone returns false if the Option is Some
func (s some[T]) IsNone() bool {
	return false
}

// IsSome returns true if the Option is Some
func (s some[T]) IsSome() bool {
	return true
}

// NewNone returns a new None
func None() none {
	return none{}
}

// NewSome returns a new Some
func Some[T any](value T) some[T] {
	return some[T]{Value: value}
}

// Unwrap returns the value of a Some or panics if the Option is None
func Unwrap[T any](o Option[T]) T {
	if o.IsNone() {
		panic("Unwrap called on None")
	}
	return o.(some[T]).Value
}

// UnwrapOr returns the value of a Some or a default value if the Option is None
func UnwrapOr[T any](o Option[T], def T) T {
	if o.IsNone() {
		return def
	}
	return o.(some[T]).Value
}
