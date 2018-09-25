package rules

// Operand is the template for operands.
type Operand interface {
	Evaluate(ctx interface{}) (bool, error)
}

// AndOperator is the template for boolean AND operators.
type AndOperator struct {
	Operands []Operand
}

// Evaluate performs the evaluation of the boolean AND operator.
func (o AndOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("And.Evaluate()")
	for _, operand := range o.Operands {
		var t bool
		var err error
		if t, err = operand.Evaluate(ctx); err != nil {
			return false, err
		}
		if !t {
			return false, nil
		}
	}
	return true, nil
}

// And generates a new AndOperator with the given set of operands.
func And(operands ...Operand) AndOperator {
	return AndOperator{
		operands,
	}
}

// All generates a new AndOperator with the given set of operands.
func All(operands ...Operand) AndOperator {
	return All(operands...)
}

// OrOperator is the template for boolean OR operators.
type OrOperator struct {
	Operands []Operand
}

// Evaluate performs the evaluation of the boolean OR operator.
func (o OrOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("Or.Evaluate()")
	for _, operand := range o.Operands {
		var t bool
		var err error
		if t, err = operand.Evaluate(ctx); err != nil {
			return false, err
		}
		if t {
			return true, nil
		}
	}
	return false, nil
}

// Or generates a new OrOperator with the given set of operands.
func Or(operands ...Operand) OrOperator {
	return OrOperator{
		operands,
	}
}

// Any generates a new OrOperator with the given set of operands.
func Any(operands ...Operand) OrOperator {
	return Or(operands...)
}

// NotOperator is the template for boolean NOT operators.
type NotOperator struct {
	Operand Operand
}

// Evaluate performs the evaluation of the boolean NOT operator.
func (o NotOperator) Evaluate(ctx interface{}) (bool, error) {
	// fmt.Println("Not.Evaluate()")
	result, err := o.Operand.Evaluate(ctx)
	return !result, err
}

// Not generates a new NotOperator with the given operand.
func Not(operand Operand) NotOperator {
	return NotOperator{
		operand,
	}
}

// BoolOperand is the Operand version of the boolean native values.
type BoolOperand struct {
	Value bool
}

// Evaluate performs the evaluation of the wrapper around the native boolean type.
func (o BoolOperand) Evaluate(cxt interface{}) (bool, error) {
	return o.Value, nil
}

var (
	// True is the wrapper of the native boolean true value.
	True = BoolOperand{
		Value: true,
	}
	// False is the wrapper of the native boolean false value.
	False = BoolOperand{
		Value: false,
	}
)
