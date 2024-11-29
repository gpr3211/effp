package effp

import (
	"fmt"
)

// Set function that works with any slice (basic types or structs)
// returns nil on empty input.
func Set[T comparable](slice []T) []T {
	uniqueMap := make(map[T]bool)
	var uniqueSlice []T
	for _, item := range slice {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueSlice = append(uniqueSlice, item)
		}
	}
	return uniqueSlice
}

// Union takes 2 sets of comparable types and returns their union.
// returns nil else
func Union[T comparable](set1, set2 []T) []T {
	combined := append(set1, set2...)
	return Set(combined)
}

// Intersection returns a slice containing the intersection of two slices.
// returns nil if there is no intersection between set1 and set2
func Intersection[T comparable](set1, set2 []T) []T {
	uniqueMap := make(map[T]bool)
	for _, item := range set1 {
		uniqueMap[item] = true
	}
	var intersectionSlice []T
	for _, item := range set2 {
		if uniqueMap[item] {
			intersectionSlice = append(intersectionSlice, item)
		}
	}
	return intersectionSlice
}

// Difference returns a slice containing elements in set1 but not in set2.
// returns nil if sets are equal
func Difference[T comparable](set1, set2 []T) []T {
	uniqueMap := make(map[T]bool)
	for _, item := range set2 {
		uniqueMap[item] = true
	}

	var differenceSlice []T
	for _, item := range set1 {
		if !uniqueMap[item] {
			differenceSlice = append(differenceSlice, item)
		}
	}
	//	fmt.Println(set1, set2)

	return differenceSlice
}

func main() {
	type Point struct {
		X int
		Y int
	}
	_ = []Point{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
		{X: 1, Y: 2}, // Duplicate
		{X: 5, Y: 6},
	}
	// Integer example
	setA := []int{1, 2, 3, 4}
	setB := []int{3, 4, 5, 6}
	setH := []int{1, 2, 3, 4}

	type Product struct {
		Name, Category string
		Price          float64
	}

	var ProductList = []Product{
		{"Kayak", "Watersports", 1},
		{"Lifejacket", "Watersports", 49.95},
		{"Soccer Ball", "Soccer", 19.50},
		{"Corner Flags", "Soccer", 34.95},
		{"Stadium", "Soccer", 79500},
		{"Stadium", "Soccer", 79500},
		{"Stadium", "Soccer", 79500},
		{"Thinking Cap", "Chess", 16},
		{"Bling-Bling King", "Chess", 1200},
		{"Bling-Bling King", "Chess", 1200},
	}
	type Config struct {
		p *Point
	}
	cfg := Config{}

	list := []Config{cfg, Config{&Point{Y: 0, X: 1}}}
	cfgMap := func(c Config) Config { return Config{&Point{c.p.Y, c.p.X * 2}} }
	newcfg, err := Map(list, cfgMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newcfg)

	salePrice := func(p Product) Product {
		return Product{p.Name, p.Category, p.Price * 2}
	}
	saleItems, err := Map(ProductList, salePrice)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := range saleItems {
			fmt.Println(saleItems[i])
		}
	}

	fmt.Println(Set(ProductList))
	fmt.Println("Union:", Union(setA, setB))
	fmt.Println("Intersection:", Intersection(setA, setB))
	fmt.Println("Difference:", Difference(setA, setB))
	// String example
	setC := []string{"apple", "banana", "cherry"}
	setD := []string{"banana", "cherry", "date"}

	fmt.Println("Union (strings):", Union(setC, setD))
	fmt.Println("Intersection (strings):", Intersection(setC, setD)) // Output: [banana cherry]
	fmt.Println("Difference (strings):", Difference(setC, setD))
	fmt.Println("Difference (zero):", Difference(setA, setH))
	fmt.Println("empty is nil")

}
