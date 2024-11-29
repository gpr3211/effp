# Functional programming Go package

## List

## FUNCTIONS

- func All[A any](input []A, pred Predicate[A]) bool
        All returns true if all elements satisfy predicate

- func Any[A any](input []A, pred Predicate[A]) bool
    Any takes in any array and returns true if predicate is true for an element.
    - type Predicate func(A any) bool

- func Difference[T comparable](set1, set2 []T) []T
    Difference returns a slice containing elements in set1 but not in set2.
    returns nil if sets are equal

- func Filter[A any](input []A, pred Predicate[A]) []A
    Filter takes as input a slice of Any and a Predicate type function to be
    applied to each item in the slice. returns a new slice containing only the
    satisfied elements

- func Intersection[T comparable](set1, set2 []T) []T
    Intersection returns a slice containing the intersection of two slices.
    returns nil if there is no intersection between set1 and set2

- func Map[A any](input []A, m MapFunc[A]) []A
    Map applies MapFunc to each element of the input slice returns a new
    modified slice with the same number of elements

- func Set[T comparable](slice []T) []T
    Set function that works with any slice (basic types or structs) returns nil
    on empty input

- func Union[T comparable](set1, set2 []T) []T
    Union takes 2 sets of comparable types and returns their union returns nil
    else


## TYPES

- type MapFunc[A any] func(A) A
    MapFunc type of Any func(A any) A

- type Predicate[A any] func(A) bool
    Predicate type takes any type of slice and a function which returns a bool
