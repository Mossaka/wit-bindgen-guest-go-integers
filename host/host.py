import sys
import wasmtime
from integers import add_integers_to_linker, Integers
from integers_export import IntegersExport

class MyIntegers(Integers):
    def a1(self, x: int) -> None:
        print(f"a1: {x}")

    def a2(self, x: int) -> None:
        print(f"a2: {x}")

def run(wasm_file: str) -> None:
    store = wasmtime.Store()
    module = wasmtime.Module.from_file(store.engine, wasm_file)
    linker = wasmtime.Linker(store.engine)
    linker.define_wasi()
    wasi = wasmtime.WasiConfig()
    wasi.inherit_stdout()
    wasi.inherit_stderr()
    store.set_wasi(wasi)

    add_integers_to_linker(linker, store, MyIntegers())
    wasm = IntegersExport(store, linker, module)

    res = wasm.res1(store)
    res = wasm.res2(store)
    print(res)

if __name__ == '__main__':
    run(sys.argv[1])
