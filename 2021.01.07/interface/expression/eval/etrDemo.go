package main

import (
	"fmt"
	"math"
	"strings"
)

//接口Expr来表示Go语言中任意的表达式
type Expr interface{
	//Eval会根据给定的Env变量返回表达式的值。因为每个表达式都必须提供这个方法，我们将它加入到Expr接口中。
	Eval(env Env) float64
	//Check报告此表达式中的错误并将其变量添加到集合中。
	Check(vars map[Var]bool) error
}

// Var 标识变量, e.g., x.
type Var string

// literal 是一个数字常量, e.g., 3.141.
type literal float64

// unary 表示一元运算符表达式, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// binary 表示二进制运算符表达式, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// call 表示函数调用表达式, 我们限制它的fn字段只能是pow，sin或者sqrt e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

//Env将变量的名字映射成对应的值：
type Env map[Var]float64

//Var类型的Eval方法对一个env变量进行查找
func (v Var) Eval(env Env) float64 {
	return env[v]
}

//literal类型的Eval方法返回它真实的值
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}
//unary的Eval方法会递归的计算它的运算对象，然后将运算符op作用到它们上
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}
//binary的Eval方法会递归的计算它的运算对象，然后将运算符op作用到它们上
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}
//最后，call的这个方法会计算对于pow，sin，或者sqrt函数的参数值，然后调用对应在math包中的函数。
func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
//检查是否为
func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}