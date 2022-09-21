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
class IntegersExport:
    instance: wasmtime.Instance
    _res: wasmtime.Func
    def __init__(self, store: wasmtime.Store, linker: wasmtime.Linker, module: wasmtime.Module):
        self.instance = linker.instantiate(store, module)
        exports = self.instance.exports(store)
        
        res = exports['res']
        assert(isinstance(res, wasmtime.Func))
        self._res = res
    def res(self, caller: wasmtime.Store) -> int:
        ret = self._res(caller)
        assert(isinstance(ret, int))
        return _clamp(ret, 0, 255)
