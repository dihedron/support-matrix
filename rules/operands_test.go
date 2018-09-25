package rules

import (
	"testing"
)

func TestAnd(t *testing.T) {

	type test struct {
		label    string
		operand  Operand
		expected bool
	}

	for _, test := range []test{
		test{
			"and (t -> t)",
			And(
				True,
			),
			true,
		},
		test{
			"and (f -> f)",
			And(
				False,
			),
			false,
		},
		test{
			"and (t,t,f -> f)",
			And(
				True,
				True,
				False,
			),
			false,
		},
		test{
			"and (t,t,t -> t)",
			And(
				True,
				True,
				True,
			),
			true,
		},

		test{
			"or (t -> t)",
			Or(
				True,
			),
			true,
		},
		test{
			"or (f -> f)",
			Or(
				False,
			),
			false,
		},
		test{
			"or (t,t,f -> f)",
			Or(
				True,
				True,
				False,
			),
			true,
		},
		test{
			"or (f,f,f -> f)",
			Or(
				False,
				False,
				False,
			),
			false,
		},
		// NOT
		test{
			"not (f -> t)",
			Not(
				False,
			),
			true,
		},
		test{
			"not (t -> f)",
			Not(
				True,
			),
			false,
		},
	} {
		t.Logf("running test %q...\n", test.label)
		actual, _ := test.operand.Evaluate(nil)
		if actual != test.expected {
			t.Fatalf("test %q: expected %t, got %t\n", test.label, test.expected, actual)
		}
	}
}
