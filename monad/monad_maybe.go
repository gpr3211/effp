package monad

import (
	"strconv"
)

// Maybe interface defines the common operations
type Maybe[A any] interface {
	Get() A
	GetOrElse(def A) A
}

// JustMaybe represents a present value
type JustMaybe[A any] struct {
	value A
}

func (j JustMaybe[A]) Get() A {
	return j.value
}

func (j JustMaybe[A]) GetOrElse(def A) A {
	return j.value
}

// NothingMaybe represents an absent value
type NothingMaybe[A any] struct{}

func (n NothingMaybe[A]) Get() A {
	return *new(A)
}

func (n NothingMaybe[A]) GetOrElse(def A) A {
	return def
}

// Constructor functions
func Just[A any](a A) JustMaybe[A] {
	return JustMaybe[A]{value: a}
}

func Nothing[A any]() Maybe[A] {
	return NothingMaybe[A]{}
}

// Functor implementation
func fmap[A, B any](m Maybe[A], mapFunc func(A) B) Maybe[B] {
	switch m.(type) {
	case JustMaybe[A]:
		j := m.(JustMaybe[A])
		return JustMaybe[B]{
			value: mapFunc(j.value),
		}
	case NothingMaybe[A]:
		return NothingMaybe[B]{}
	default:
		panic("unknown type")
	}
}

// Example types and helper functions
type User struct {
	Name string
	Age  int
}

// Safe map access
func getFromMap[K comparable, V any](m map[K]V, key K) Maybe[V] {
	if value, ok := m[key]; ok {
		return Just(value)
	}
	return Nothing[V]()
}

// Example functions
func getUserByID(id int) Maybe[User] {
	users := map[int]User{
		1: {"Alice", 30},
		2: {"Bob", 25},
	}
	if user, exists := users[id]; exists {
		return Just(user)
	}
	return Nothing[User]()
}

func parseNumber(s string) Maybe[int] {
	if num, err := strconv.Atoi(s); err == nil {
		return Just(num)
	}
	return Nothing[int]()
}

func doubleAge(u User) int {
	return u.Age * 2
}

// Example using fromNullable
func fromNullable[A any](ptr *A) Maybe[A] {
	if ptr == nil {
		return Nothing[A]()
	}
	return Just(*ptr)
}

/*
func main() {
	// Example 1: Basic Maybe usage
	name := "John"
	maybeNamePtr := fromNullable(&name)
	fmt.Printf("Name: %v\n", maybeNamePtr.GetOrElse("Unknown"))

	var nilName *string
	maybeNilPtr := fromNullable(nilName)
	fmt.Printf("Nil name: %v\n", maybeNilPtr.GetOrElse("Unknown"))

	// Example 2: Safe user lookup
	user1 := getUserByID(1)
	user2 := getUserByID(999)

	fmt.Printf("User 1: %v\n", user1.GetOrElse(User{Name: "Unknown", Age: 0}))
	fmt.Printf("User 2: %v\n", user2.GetOrElse(User{Name: "Unknown", Age: 0}))

	// Example 3: Chain of transformations
	result := fmap(getUserByID(1), doubleAge)
	fmt.Printf("Doubled age: %v\n", result.GetOrElse(-1))

	// Example 4: Safe number parsing and transformation
	parseAndDouble := func(s string) Maybe[int] {
		return fmap(parseNumber(s), func(n int) int {
			return n * 2
		})
	}

	num1 := parseAndDouble("10")
	num2 := parseAndDouble("invalid")

	fmt.Printf("Parsed and doubled 10: %v\n", num1.GetOrElse(-1))
	fmt.Printf("Parsed and doubled invalid: %v\n", num2.GetOrElse(-1))

	// Example 5: Map operations
	data := map[string]int{
		"foo": 42,
		"bar": 84,
	}

	lookup1 := getFromMap(data, "foo")
	lookup2 := getFromMap(data, "nonexistent")

	fmt.Printf("Lookup foo: %v\n", lookup1.GetOrElse(-1))
	fmt.Printf("Lookup nonexistent: %v\n", lookup2.GetOrElse(-1))

	// Example 6: Complex chaining
	complexOperation := fmap(
		fmap(parseNumber("30"), func(age int) User {
			return User{Name: "Dynamic", Age: age}
		}),
		doubleAge,
	)
	fmt.Printf("Complex chain result: %v\n", complexOperation.GetOrElse(-1))
}
*/
