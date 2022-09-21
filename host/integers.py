from abc import abstractmethod
from typing import Any
import wasmtime

try:
    from typing import Protocol
except ImportError:
    class Protocol: # type: ignore
        pass


def _clamp(i: int, min: int, max: int) -> int:
    if i < min or i > max:
        raise OverflowError(f'must be between {min} and {max}')
    return i
class Integers(Protocol):
    @abstractmethod
    def a1(self, x: int) -> None:
        raise NotImplementedError
    @abstractmethod
    def a2(self, x: int) -> None:
        raise NotImplementedError

def add_integers_to_linker(linker: wasmtime.Linker, store: wasmtime.Store, host: Integers) -> None:
    ty = wasmtime.FuncType([wasmtime.ValType.i32()], [])
    def a1(caller: wasmtime.Caller, arg0: int) -> None:
        host.a1(_clamp(arg0, 0, 255))
    linker.define('integers', 'a1', wasmtime.Func(store, ty, a1, access_caller = True))
    ty = wasmtime.FuncType([wasmtime.ValType.i32()], [])
    def a2(caller: wasmtime.Caller, arg0: int) -> None:
        host.a2(_clamp(arg0, -128, 127))
    linker.define('integers', 'a2', wasmtime.Func(store, ty, a2, access_caller = True))
