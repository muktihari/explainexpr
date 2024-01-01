build:
	GOOS=js GOARCH=wasm go build -o ../main.wasm

build-tiny:
	tinygo build -o ../main.wasm -target wasm -no-debug -panic=trap main.go