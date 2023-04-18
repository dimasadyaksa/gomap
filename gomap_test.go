package gomap_test

import (
	"fmt"
	"testing"

	"github.com/dimasadyaksa/gomap"
)

func TestKeys(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := gomap.Keys(m)

	if len(k) != 3 {
		t.Errorf("expect 3 keys; got %d keys", len(k))
	}

	for _, key := range k {
		if _, ok := m[key]; !ok {
			t.Errorf("expect key: %s", key)
		}
	}
}

func TestValues(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	v := gomap.Values(m)

	if len(v) != 3 {
		t.Errorf("expect 3 values; got %d values", len(v))
	}

	for _, value := range v {
		if m[value] != value {
			t.Errorf("expect value: %d; got value: %d", value, m[value])
		}
	}
}

func TestFind(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	v, ok := gomap.Find(m, func(key string, value int) bool {
		return key == "two"
	})

	if !ok {
		t.Error("expect to find the key")
	}

	if v != 2 {
		t.Errorf("expect value: 2; got value: %d", v)
	}

	v, ok = gomap.Find(m, func(key string, value int) bool {
		return key == "four"
	})

	if ok {
		t.Error("expect not to find the key")
	}

	if v != 0 {
		t.Errorf("expect value: 0; got value: %d", v)
	}
}

func TestReduce(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	r := gomap.Reduce(m, 0, func(r int, key string, value int) int {
		return r + value
	})

	if r != 6 {
		t.Errorf("expect value: 6; got value: %d", r)
	}

	str := gomap.Reduce(m, "", func(r string, key string, value int) string {
		return r + key
	})

	if len(str) != len("onetwothree") {
		t.Errorf("expect value: onetwothree; got value: %s", str)
	}
}

func TestSlice(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	v := gomap.Values(m)

	if len(v) != 3 {
		t.Errorf("expect 3 values; got %d values", len(v))
	}

	for _, value := range v {
		if m[value] != value {
			t.Errorf("expect value: %d; got value: %d", value, m[value])
		}
	}
}

func TestFilter(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	v := gomap.Filter(m, func(key int, value int) bool {
		return value > 1
	})

	if len(v) != 2 {
		t.Errorf("expect 2 values; got %d values", len(v))
	}

	for _, value := range v {
		if m[value] != value {
			t.Errorf("expect value: %d; got value: %d", value, m[value])
		}
	}
}

func TestCombine(t *testing.T) {
	m1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	m2 := map[int]int{
		4: 4,
		5: 5,
		6: 6,
	}

	c := gomap.Combine(m1, m2)

	if len(c) != 6 {
		t.Errorf("expect 6 values; got %d values", len(c))
	}

	for key, value := range c {
		if c[key] != value {
			t.Errorf("expect value: %d; got value: %d", value, c[key])
		}
	}
}

func TestCombineWithOverwrite(t *testing.T) {
	m1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	m2 := map[int]int{
		3: 4,
		4: 5,
		5: 6,
	}

	c := gomap.Combine(m1, m2)

	if len(c) != 5 {
		t.Errorf("expect 5 values; got %d values", len(c))
	}

	if c[3] != 4 {
		t.Errorf("expect value: 4; got value: %d", c[3])
	}
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}

	v := gomap.Map(arr, func(value int) string {
		return fmt.Sprint(value + 1)
	})

	if len(v) != 3 {
		t.Errorf("expect 3 values; got %d values", len(v))
	}

	for _, value := range v {
		if v[fmt.Sprint(value)] != value {
			t.Errorf("expect value: %d; got value: %d", value, v[fmt.Sprint(value)])
		}
	}
}

func TestEvery(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if !gomap.Every(m, func(key string, value int) bool {
		return value > 0
	}) {
		t.Error("expect true")
	}

	if gomap.Every(m, func(key string, value int) bool {
		return value > 1
	}) {
		t.Error("expect false")
	}
}

func TestSome(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if !gomap.Some(m, func(key string, value int) bool {
		return value > 1
	}) {
		t.Error("expect true")
	}

	if gomap.Some(m, func(key string, value int) bool {
		return value > 3
	}) {
		t.Error("expect false")
	}
}

func TestEqual(t *testing.T) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if !gomap.Equal(m1, m2) {
		t.Error("expect true")
	}

	m2["three"] = 4

	if gomap.Equal(m1, m2) {
		t.Error("expect false")
	}
}

func TestEqualWithDifferentLength(t *testing.T) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	if gomap.Equal(m1, m2) {
		t.Error("expect false")
	}
}

func TestEqualFunc(t *testing.T) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	if !gomap.EqualFunc(m1, m2, func(a int, b int) bool {
		return a == b
	}) {
		t.Error("expect true")
	}

	m2["three"] = 4

	if gomap.EqualFunc(m1, m2, func(a int, b int) bool {
		return a == b
	}) {
		t.Error("expect false")
	}
}

func TestEqualFuncWithDifferentLength(t *testing.T) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	if gomap.EqualFunc(m1, m2, func(a int, b int) bool {
		return a == b
	}) {
		t.Error("expect false")
	}
}

func TestClear(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	gomap.Clear(m)

	if len(m) != 0 {
		t.Errorf("expect 0 values; got %d values", len(m))
	}
}

func BenchmarkKeys(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i := 0; i < b.N; i++ {
		gomap.Keys(m)
	}
}

func BenchmarkValues(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i := 0; i < b.N; i++ {
		gomap.Values(m)
	}
}

func BenchmarkFind(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	predicate := func(key string, value int) bool { return key == "two" }

	for i := 0; i < b.N; i++ {
		gomap.Find(m, predicate)
	}
}

func BenchmarkReduce(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	reducer := func(r int, key string, value int) int { return r + value }

	for i := 0; i < b.N; i++ {
		gomap.Reduce(m, 0, reducer)
	}
}

func BenchmarkFilter(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	predicate := func(key string, value int) bool { return value > 1 }

	for i := 0; i < b.N; i++ {
		gomap.Filter(m, predicate)
	}
}

func BenchmarkCombine(b *testing.B) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
	}

	for i := 0; i < b.N; i++ {
		gomap.Combine(m1, m2)
	}
}

func BenchmarkMap(b *testing.B) {
	slice := []int{1, 2, 3}
	keyFunc := func(value int) int { return value }

	for i := 0; i < b.N; i++ {
		gomap.Map(slice, keyFunc)
	}
}

func BenchmarkEvery(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	predicate := func(key string, value int) bool { return value > 0 }

	for i := 0; i < b.N; i++ {
		gomap.Every(m, predicate)
	}
}

func BenchmarkSome(b *testing.B) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	predicate := func(key string, value int) bool { return value > 1 }

	for i := 0; i < b.N; i++ {
		gomap.Some(m, predicate)
	}
}

func BenchmarkEqual(b *testing.B) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i := 0; i < b.N; i++ {
		gomap.Equal(m1, m2)
	}
}

func BenchmarkEqualFunc(b *testing.B) {
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	eqFunc := func(a int, b int) bool { return a == b }

	for i := 0; i < b.N; i++ {
		gomap.EqualFunc(m1, m2, eqFunc)
	}
}

func BenchmarkClear(b *testing.B) {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	for i := 0; i < b.N; i++ {
		gomap.Clear(m)

		m[1] = 1
		m[2] = 2
		m[3] = 3
	}
}
