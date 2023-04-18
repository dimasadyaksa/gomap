package gomap

type Predicate[K comparable, V any] func(K, V) bool

// Keys returns a slice of keys from a map
// The order of the values is not guaranteed
func Keys[K comparable, V any](m map[K]V) []K {
	return Reduce(m, make([]K, 0, len(m)), keyReducer[K, V])
}

// Values returns a slice of values from a map
// The order of the values is not guaranteed
func Values[K comparable, V any](m map[K]V) []V {
	return Reduce(m, make([]V, 0, len(m)), valueReducer[K, V])
}

func valueReducer[K, V any](s []V, _ K, v V) []V                   { return append(s, v) }
func keyReducer[K, V any](s []K, k K, _ V) []K                     { return append(s, k) }
func copyReducer[K comparable, V any](r map[K]V, k K, v V) map[K]V { r[k] = v; return r }
func conditionalValueReducer[K comparable, V any](p Predicate[K, V]) func([]V, K, V) []V {
	return func(s []V, k K, v V) []V {
		if p(k, v) {
			return append(s, v)
		}
		return s
	}
}

// Find returns the first value that matches the predicate
func Find[K comparable, V any](m map[K]V, predicate Predicate[K, V]) (V, bool) {
	var v V

	for key, value := range m {
		if predicate(key, value) {
			return value, true
		}
	}

	return v, false
}

// Reduce returns a single value by applying the reducer function to each value in the map
// The order of the values is not guaranteed
func Reduce[K comparable, V any, R any](m map[K]V, initial R, reducer func(R, K, V) R) R {
	reduced := initial
	for k, v := range m {
		reduced = reducer(reduced, k, v)
	}

	return reduced
}

// Filter returns a map of values that match the predicate
func Filter[K comparable, V any](m map[K]V, predicate Predicate[K, V]) map[K]V {
	return Reduce(m, make(map[K]V, len(m)), func(r map[K]V, k K, v V) map[K]V {
		if predicate(k, v) {
			r[k] = v
		}
		return r
	})
}

// Combine combines two maps into a single map. If a key exists in both maps, the value from the second map will be used.
func Combine[K comparable, V any](m1, m2 map[K]V) map[K]V {
	return combine(m1, m2, copyReducer[K, V])
}

func combine[K comparable, V any](m1, m2 map[K]V, reducer func(r map[K]V, k K, v V) map[K]V) map[K]V {
	m := make(map[K]V, len(m1)+len(m2))
	return Reduce(m2, Reduce(m1, m, reducer), reducer)
}

// Map returns a map of values from a slice
func Map[K comparable, V any](s []V, key func(idx int) K) map[K]V {
	m := make(map[K]V, len(s))

	for i, value := range s {
		m[key(i)] = value
	}

	return m
}

// Every returns true if every value in the map matches the predicate
func Every[K comparable, V any](m map[K]V, predicate Predicate[K, V]) bool {
	return Reduce(m, true, func(b bool, k K, v V) bool { return b && predicate(k, v) })
}

// Some returns true if any value in the map matches the predicate
func Some[K comparable, V any](m map[K]V, predicate Predicate[K, V]) bool {
	_, ok := Find(m, predicate)
	return ok
}

// Equal returns true if the two maps are equal
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	return EqualFunc[K, V](m1, m2, eq[V])
}

// EqualFunc returns true if the two maps are equal using the equal function
func EqualFunc[K comparable, V any](m1, m2 map[K]V, equal func(V, V) bool) bool {
	return len(m1) == len(m2) && Every(m1, func(k K, v V) bool { v2, ok := m2[k]; return ok && equal(v, v2) })
}

func eq[T comparable](a, b T) bool { return a == b }

// Clear removes all values from the map
func Clear[K comparable, V any](m map[K]V) {
	Reduce(m, m, func(r map[K]V, k K, v V) map[K]V { delete(r, k); return r })
}

// FindAll returns a slice of values that match the predicate
func FindAll[K comparable, V any](m map[K]V, predicate Predicate[K, V]) []V {
	return Reduce(m, make([]V, 0, len(m)), conditionalValueReducer[K, V](predicate))
}

// Intersect returns a map of values that exist in both maps
func Intersect[K comparable, V any](m1, m2 map[K]V) map[K]V {
	return Filter(m1, func(k K, v V) bool { _, ok := m2[k]; return ok })
}
