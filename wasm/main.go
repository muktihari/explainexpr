package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall/js"

	"github.com/muktihari/expr"
	"github.com/muktihari/expr/explain"
)

type Result struct {
	Value any    `json:"value"`
	Err   string `json:"err"`
}

func Explain() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if !isInputValid(args) {
			return encode(&Result{Err: "input is empty"})
		}

		s := args[0].String()
		steps, err := explain.Explain(s)
		if err != nil {
			return encode(&Result{Err: fmt.Sprintf("%s", err)})
		}

		return encode(&Result{Value: steps})
	})
}

func Evaluate() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if !isInputValid(args) {
			return encode(&Result{Err: "input is empty"})
		}

		s := args[0].String()
		v, err := expr.Any(s)
		if err != nil {
			return encode(&Result{Err: fmt.Sprintf("%s", err)})
		}

		return encode(&Result{Value: v})
	})
}

func encode(res *Result) string {
	b, _ := json.Marshal(res)
	return string(b)
}

func isInputValid(args []js.Value) bool {
	if len(args) == 0 {
		return false
	}

	s := args[0].String()
	if s == "" {
		return false
	}

	return true
}

func main() {
	fmt.Printf("Hello Web Assembly from Go!\n") // printed on browser console
	ch := make(chan os.Signal, 1)
	js.Global().Set("explain", Explain())
	js.Global().Set("evaluate", Evaluate())
	<-ch // never exit
}
