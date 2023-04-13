package gomap

// Keys returns a slice of keys from a map
// The order of the values is not guaranteed
func Keys[K comparable, V any](m map[K]V) []K {
	return reduce(m, maker[K, V, K], func(ks []K, k K, v V) []K { return append(ks, k) })
}

// Values returns a slice of values from a map
// The order of the values is not guaranteed
func Values[K comparable, V any](m map[K]V) []V {
	return reduce(m, maker[K, V, V], func(vs []V, k K, v V) []V { return append(vs, v) })
}

func maker[K comparable, V any, S any](m map[K]V) []S { return make([]S, 0, len(m)) }

func reduce[K comparable, V any, R any](m map[K]V, initial func(map[K]V) R, reducer func(R, K, V) R) R {
	r := initial(m)
	for key, value := range m {
		r = reducer(r, key, value)
	}
	return r
}

// Find returns the first value that matches the predicate
func Find[K comparable, V any](m map[K]V, predicate func(K, V) bool) (V, bool) {
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
	return reduce(m, func(map[K]V) R { return initial }, reducer)
}

// Filter returns a map of values that match the predicate
func Filter[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	f := make(map[K]V, len(m))

	for key, value := range m {
		if predicate(key, value) {
			f[key] = value
		}
	}

	return f
}

// Combine combines two maps into a single map. If a key exists in both maps, the value from the second map will be used.
func Combine[K comparable, V any](m1, m2 map[K]V) map[K]V {
	m := make(map[K]V, len(m1)+len(m2))

	for key, value := range m1 {
		m[key] = value
	}

	for key, value := range m2 {
		m[key] = value
	}

	return m
}

// Map returns a map of values from a slice
func Map[K comparable, V any](s []V, key func(V) K) map[K]V {
	m := make(map[K]V, len(s))

	for _, value := range s {
		m[key(value)] = value
	}

	return m
}

// Every returns true if every value in the map matches the predicate
func Every[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	return reduce(m, func(map[K]V) bool { return true }, func(b bool, k K, v V) bool { return b && predicate(k, v) })
}

// Some returns true if any value in the map matches the predicate
func Some[K comparable, V any](m map[K]V, predicate func(K, V) bool) bool {
	for key, value := range m {
		if predicate(key, value) {
			return true
		}
	}

	return false
}

// Equal returns true if the two maps are equal
func Equal[K, V comparable](m1, m2 map[K]V) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}

	return true
}

// EqualFunc returns true if the two maps are equal using the equal function
func EqualFunc[K comparable, V any](m1, m2 map[K]V, equal func(V, V) bool) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !equal(v1, v2) {
			return false
		}
	}

	return true
}
