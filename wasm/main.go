package main

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/muktihari/expr"
	expln "github.com/muktihari/expr/exp/explain"
)

// linierMemory is a WASM Linier Memory.
//
// ref: https://wasmbyexample.dev/examples/webassembly-linear-memory/webassembly-linear-memory.go.en-us.html
var linierMemory = []byte{}

func main() {
	fmt.Printf("Explain Expr Web Assembly instantiated!\n") // printed on browser console
}

type Result struct {
	Value any    `json:"value"`
	Err   string `json:"err"`
}

// encode puts result into liner memory then return the address.
func encode(res *Result) uintptr {
	linierMemory, _ = json.Marshal(res)
	return uintptr(unsafe.Pointer(&linierMemory[0]))
}

// lmalloc stands for Linier Memory Allocation, it allocate size amount of memory that can be filled from Javascript.
//
//go:export lmalloc
func lmalloc(size int) uintptr {
	linierMemory = make([]byte, size)
	return uintptr(unsafe.Pointer(&linierMemory[0]))
}

// size returns the bytes size of encoded result in linierMemory
//
//go:export size
func size() int { return len(linierMemory) }

// explain explain given expression
//
//go:export explain
func explain() uintptr {
	if len(linierMemory) == 0 {
		return encode(&Result{Err: "input is empty"})
	}
	steps, err := expln.Explain(string(linierMemory))
	if err != nil {
		return encode(&Result{Err: fmt.Sprintf("%s", err)})
	}
	return encode(&Result{Value: steps})
}

// explain evaluate given expression
//
//go:export evaluate
func evaluate() uintptr {
	if len(linierMemory) == 0 {
		return encode(&Result{Err: "input is empty"})
	}
	value, err := expr.Any(string(linierMemory))
	if err != nil {
		return encode(&Result{Err: fmt.Sprintf("%s", err)})
	}
	return encode(&Result{Value: value})
}
