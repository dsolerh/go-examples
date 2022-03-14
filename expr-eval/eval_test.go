package expreval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	testCases := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, tC := range testCases {
		// Print expr only when it changes.
		if tC.expr != prevExpr {
			fmt.Printf("\nExpression: %s", tC.expr)
			prevExpr = tC.expr
		}
		expr, err := Parse(tC.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(tC.env))
		fmt.Printf("\t%v => %s\n", tC.env, got)
		if got != tC.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", tC.expr, tC.env, got, tC.want)
		}
	}
}

func TestCoverage(t *testing.T) {
	testCases := []struct {
		input string
		env   Env
		want  string // expected error from Parse/Check or result from Eval
	}{
		{"x % 2", nil, "unexpected '%'"},
		{"!true", nil, "unexpected '!'"},
		{"log(10)", nil, `unknown function "log"`},
		{"sqrt(1, 2)", nil, "call to sqrt has 2 args, want 1"},
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"+x * -x", Env{"x": -4}, "-16"},
	}
	for _, tC := range testCases {
		expr, err := Parse(tC.input)
		if err == nil {
			err = expr.Check(map[Var]bool{})
		}
		if err != nil {
			if err.Error() != tC.want {
				t.Errorf("%s: got %q, want %q", tC.input, err, tC.want)
			}
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(tC.env))
		if got != tC.want {
			t.Errorf("%s: %v => %s, want %s", tC.input, tC.env, got, tC.want)
		}
	}
}
