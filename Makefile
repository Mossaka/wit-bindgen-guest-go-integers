.PHONY: build
build:
	tinygo build -wasm-abi=generic -target=wasi -gc=leaking -no-debug -o main.wasm main.go