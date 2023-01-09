import sys
from tarfile import StreamError
import wasmtime
from host import Integers, IntegersImports, WasiStream, imports, wasi_filesystem, wasi_logging, wasi_poll
from host.imports.wasi_filesystem import Descriptor, Errno, Filesize, Size
from host.imports.wasi_logging import Level
from host.types import Result

class MyIntegers(imports.Imports):
    def a1(self, x: int) -> None:
        print(f"a1: {x}")

    def a2(self, x: int) -> None:
        print(f"a2: {x}")

class Logging(wasi_logging.WasiLogging):
    def log(self, level: Level, context: str, message: str) -> None:
        print(f"{message}")

class Filesystem(wasi_filesystem.WasiFilesystem):
    def write_via_stream(self, fd: Descriptor, offset: Filesize) -> Result[WasiStream, Errno]:
        raise NotImplementedError

class Poll(wasi_poll.WasiPoll):
    def write_stream(self, stream: WasiStream, buf: bytes) -> Result[Size, StreamError]:
        raise NotImplementedError

def run() -> None:
    store = wasmtime.Store()
    wasm = Integers(store, IntegersImports(MyIntegers(), Logging(), Filesystem(), Poll()))

    res = wasm.res(store)

    print(res)

if __name__ == '__main__':
    run()
