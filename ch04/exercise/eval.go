package eval

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// BinOp is binary operator like +-*/
type BinOp func(int, int) int

// OpDef is a mapping of operator(+-*/) to BinOp
// Describes how operator should work
type OpDef map[string]BinOp

// PriorityDef defines priority of operator
type PriorityDef map[string]Set

// Evaluator reads expression and returns result of expr.
type Evaluator struct {
	op OpDef
	pr PriorityDef
}

var rxTpl = regexp.MustCompile(`({[^}]+})`)

// NewEvaluator returns evaluator function
func NewEvaluator(ops OpDef, prs PriorityDef) Evaluator {
	return Evaluator{
		op: ops,
		pr: prs,
	}
}

// Parse reads given template and returns parsed output
func (e Evaluator) Parse(template string) string {
	return rxTpl.ReplaceAllStringFunc(template, func(expr string) string {
		expr = strings.Trim(expr, "{ }")
		result, err := e.Eval(expr)
		if err != nil {
			return ""
		}

		ret := fmt.Sprintf("%d", result)
		return ret
	})
}

// Eval read expression and returns evaluated result with error
func (e Evaluator) Eval(expr string) (int, error) {
	ops := NewStack("(")
	values := NewStack()
	reduce := func(nextOp string) error {
		for !ops.Empty() {
			tmp, err := ops.Peek()
			if err != nil {
				return err
			}
			op := tmp.(string)

			if _, higher := e.pr[nextOp][op]; nextOp != ")" && !higher {
				return nil
			}
			ops.Pop()

			if op == "(" {
				return nil
			}

			tmp, err = values.Pop()
			if err != nil {
				return err
			}
			b := tmp.(int)

			tmp, err = values.Pop()
			if err != nil {
				return err
			}
			a := tmp.(int)

			if f := e.op[op]; f != nil {
				values.Push(f(a, b))
			}
		}
		return nil
	}

	r := bufio.NewScanner(strings.NewReader(expr))
	r.Split(bufio.ScanWords)
	for r.Scan() {
		token := r.Text()
		if token == "(" {
			ops.Push(token)
		} else if _, ok := e.pr[token]; ok {
			reduce(token)
			ops.Push(token)
		} else if token == ")" {
			reduce(token)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return -1, err
			}
			values.Push(num)
		}
	}
	reduce(")")
	ret, err := values.Peek()
	return ret.(int), err
}
