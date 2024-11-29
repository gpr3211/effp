package effp

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// Product struct for testing
type Product struct {
	Name     string
	Category string
	Price    float64
}
type Point struct {
	X, Y int
}

// Custom type for testing
type CustomID int

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
func TestSetWithVariousTypes(t *testing.T) {
	t.Run("Integer slice", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{
				name:     "Set integers",
				input:    []int{1, 2, 2, 3, 4, 4, 5},
				expected: []int{1, 2, 3, 4, 5},
			},
			{
				name:     "All unique integers",
				input:    []int{6, 7, 8, 9},
				expected: []int{6, 7, 8, 9},
			},
			{
				name:     "Empty integer slice",
				input:    []int{},
				expected: []int(nil),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Set(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("String slice", func(t *testing.T) {
		tests := []struct {
			name     string
			input    []string
			expected []string
		}{
			{
				name:     "Set strings",
				input:    []string{"apple", "banana", "apple", "cherry", "banana"},
				expected: []string{"apple", "banana", "cherry"},
			},
			{
				name:     "All unique strings",
				input:    []string{"x", "y", "z"},
				expected: []string{"x", "y", "z"},
			},
			{
				name:     "Empty string slice",
				input:    []string{},
				expected: []string(nil),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Set(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})

	t.Run("Custom struct slice", func(t *testing.T) {

		tests := []struct {
			name     string
			input    []Point
			expected []Point
		}{
			{
				name: "Set points",
				input: []Point{
					{X: 1, Y: 2},
					{X: 3, Y: 4},
					{X: 1, Y: 2}, // Duplicate
					{X: 5, Y: 6},
				},
				expected: []Point{
					{X: 1, Y: 2},
					{X: 3, Y: 4},
					{X: 5, Y: 6},
				},
			},
			{
				name: "All unique points",
				input: []Point{
					{X: 7, Y: 8},
					{X: 9, Y: 10},
				},
				expected: []Point{
					{X: 7, Y: 8},
					{X: 9, Y: 10},
				},
			},
			{
				name:     "Empty point slice",
				input:    []Point{},
				expected: []Point(nil), // == to nil
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Set(tt.input)
				assert.Equal(t, tt.expected, result)
			})
		}
	})
}
func TestSet(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Integer slice with duplicates",
			input:    []int{1, 2, 2, 3, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Empty integer slice",
			input:    []int{},
			expected: []int(nil),
		},
		{
			name:     "String slice with duplicates",
			input:    []string{"apple", "banana", "apple", "cherry", "banana"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "Custom type with duplicates",
			input:    []CustomID{1, 1, 2, 3, 3},
			expected: []CustomID{1, 2, 3},
		},
		{
			name: "Product struct slice with duplicates",
			input: []Product{
				{"Kayak", "Watersports", 279},
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
				{"Soccer Ball", "Soccer", 19.50},
			},
			expected: []Product{
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch input := tt.input.(type) {
			case []int:
				result := Set(input)
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Set() = %v, want %v", result, expected)
				}
			case []string:
				result := Set(input)
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Set() = %v, want %v", result, expected)
				}
			case []CustomID:
				result := Set(input)
				expected := tt.expected.([]CustomID)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Set() = %v, want %v", result, expected)
				}
			case []Product:
				result := Set(input)
				expected := tt.expected.([]Product)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Set() = %v, want %v", result, expected)
				}
			}
		})
	}

}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		set1     interface{}
		set2     interface{}
		expected interface{}
	}{
		{
			name:     "Integer slices",
			set1:     []int{1, 2, 3},
			set2:     []int{3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "String slices",
			set1:     []string{"apple", "banana"},
			set2:     []string{"banana", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name: "Product struct slices",
			set1: []Product{
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
			},
			set2: []Product{
				{"Soccer Ball", "Soccer", 19.50},
				{"Chess Board", "Chess", 50.00},
			},
			expected: []Product{
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
				{"Chess Board", "Chess", 50.00},
			},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch set1 := tt.set1.(type) {
			case []int:
				result := Union(set1, tt.set2.([]int))
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Union() = %v, want %v", result, expected)
				}
			case []string:
				result := Union(set1, tt.set2.([]string))
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Union() = %v, want %v", result, expected)
				}
			case []Product:
				result := Union(set1, tt.set2.([]Product))
				expected := tt.expected.([]Product)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Union() = %v, want %v", result, expected)
				}
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		set1     interface{}
		set2     interface{}
		expected interface{}
	}{
		{
			name:     "Integer slices",
			set1:     []int{1, 2, 3, 4},
			set2:     []int{3, 4, 5, 6},
			expected: []int{3, 4},
		},
		{
			name:     "String slices",
			set1:     []string{"apple", "banana", "cherry"},
			set2:     []string{"banana", "cherry", "date"},
			expected: []string{"banana", "cherry"},
		},
		{
			name: "Product struct slices",
			set1: []Product{
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
			},
			set2: []Product{
				{"Soccer Ball", "Soccer", 19.50},
				{"Chess Board", "Chess", 50.00},
			},
			expected: []Product{
				{"Soccer Ball", "Soccer", 19.50},
			},
		},
		{
			name:     "Empty intersection",
			set1:     []int{1, 2, 3},
			set2:     []int{4, 5, 6},
			expected: []int(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch set1 := tt.set1.(type) {
			case []int:
				result := Intersection(set1, tt.set2.([]int))
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Intersection() = %v, want %v", result, expected)
				}
			case []string:
				result := Intersection(set1, tt.set2.([]string))
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Intersection() = %v, want %v", result, expected)
				}
			case []Product:
				result := Intersection(set1, tt.set2.([]Product))
				expected := tt.expected.([]Product)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Intersection() = %v, want %v", result, expected)
				}
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		set1     interface{}
		set2     interface{}
		expected interface{}
	}{
		{
			name:     "Integer slices",
			set1:     []int{1, 2, 3, 4},
			set2:     []int{3, 4, 5, 6},
			expected: []int{1, 2},
		},
		{
			name:     "String slices",
			set1:     []string{"apple", "banana", "cherry"},
			set2:     []string{"banana", "cherry", "date"},
			expected: []string{"apple"},
		},
		{
			name: "Product struct slices",
			set1: []Product{
				{"Kayak", "Watersports", 279},
				{"Soccer Ball", "Soccer", 19.50},
			},
			set2: []Product{
				{"Soccer Ball", "Soccer", 19.50},
				{"Chess Board", "Chess", 50.00},
			},
			expected: []Product{
				{"Kayak", "Watersports", 279},
			},
		},
		{
			name:     "Empty difference",
			set1:     []int{1, 2, 3},
			set2:     []int{1, 2, 3},
			expected: []int(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch set1 := tt.set1.(type) {
			case []int:
				result := Difference(set1, tt.set2.([]int))
				expected := tt.expected.([]int)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Difference() = %v, want %v", result, expected)
				}
			case []string:
				result := Difference(set1, tt.set2.([]string))
				expected := tt.expected.([]string)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Difference() = %v, want %v", result, expected)
				}
			case []Product:
				result := Difference(set1, tt.set2.([]Product))
				expected := tt.expected.([]Product)
				if !reflect.DeepEqual(result, expected) {
					t.Errorf("Difference() = %v, want %v", result, expected)
				}
			}
		})
	}
}
