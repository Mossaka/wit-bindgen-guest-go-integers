.PHONY: build
build:
	tinygo build -wasm-abi=generic -target=wasi -gc=leaking -no-debug -o integers.wasm main.go
	wasm-tools component new integers.wasm -o integers.wasm --adapt wasi_snapshot_preview1.wasm --wit integers.wit
	python -m wasmtime.bindgen integers.wasm --out-dir host

test: build
	python test.py