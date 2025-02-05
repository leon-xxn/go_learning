package tag

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagReflect(t *testing.T) {
	type Author struct {
		Name        string `json:"name"`
		Publication string `json:"publication,omitempty"`
	}
	tag := reflect.TypeOf(Author{})
	for i := 0; i < tag.NumField(); i++ {
		name := tag.Field(i).Name
		s, _ := tag.FieldByName(name)
		fmt.Println(name, s.Tag)
	}
}

func BenchmarkTestEq(b *testing.B) {
	a := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = testEq(a, c)
	}
}

func BenchmarkDeepEqual(b *testing.B) {
	a := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	c := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 21}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = reflect.DeepEqual(a, c)
	}
}

func testEq(a, b []int32) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
