package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall/js"

	"github.com/muktihari/expr/explain"
)

type Result struct {
	Steps []explain.Step `json:"steps"`
	Err   string         `json:"err"`
}

func Explain() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return encode(&Result{Err: "input is empty"})
		}

		s := args[0].String()
		if s == "" {
			return encode(&Result{Err: "input is empty"})
		}

		steps, err := explain.Explain(s)
		if err != nil {
			return encode(&Result{Err: fmt.Sprintf("%s", err)})
		}

		return encode(&Result{Steps: steps})
	})
}

func encode(res *Result) string {
	b, _ := json.Marshal(res)
	fmt.Printf("%s", b)
	return string(b)
}

func main() {
	fmt.Printf("hello web assembly from go!\n")
	ch := make(chan os.Signal, 1)
	js.Global().Set("explain", Explain())
	<-ch // never exit
}
