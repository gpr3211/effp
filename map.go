package effp

import "github.com/gpr3211/effp/assert"

// MapFunc type of Any func(A any) A
type MapFunc[A any] func(A) A

// Map applies MapFunc to each element of the input slice
// returns a new modified slice with the same number of elements
func Map[A any](input []A, m MapFunc[A]) ([]A, error) {
	output := make([]A, len(input))
	for i, element := range input {
		err := assert.AssertNotNil(element)
		if err != nil {
			return nil, err
		}
		output[i] = m(element)
	}
	return output, nil
}
