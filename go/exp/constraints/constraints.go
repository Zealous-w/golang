package constraints

import "golang.org/x/exp/constraints"

// Number is a constraint that permits any number type: any type
// that supports the operators < <= >= > - + * /.
// If future releases of Go add new number types,
// this constraint will be modified to include them.
type Number interface {
	constraints.Integer | constraints.Float
}
