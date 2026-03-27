package semver

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	v100, _ := Parse("1.0.0")
	v010, _ := Parse("0.1.0")
	v001b, _ := Parse("0.0.1+b")
	v001, _ := Parse("0.0.1")
	v001a, _ := Parse("0.0.1+a")
	versions := []Version{v010, v100, v001b, v001, v001a}
	Sort(versions)

	correct := []Version{v001b, v001, v001a, v010, v100}
	if !reflect.DeepEqual(versions, correct) {
		t.Fatalf("Sort returned wrong order: %s", versions)
	}
}

func BenchmarkSort(b *testing.B) {
	v100, _ := Parse("1.0.0")
	v010, _ := Parse("0.1.0")
	v001b, _ := Parse("0.0.1+b")
	v001, _ := Parse("0.0.1")
	v001a, _ := Parse("0.0.1+a")
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Sort([]Version{v010, v100, v001b, v001, v001a})
	}
}
