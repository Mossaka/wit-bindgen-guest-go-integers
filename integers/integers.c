#include <stdlib.h>
#include <integers.h>

__attribute__((weak, export_name("canonical_abi_realloc")))
void *canonical_abi_realloc(
void *ptr,
size_t orig_size,
size_t org_align,
size_t new_size
) {
  void *ret = realloc(ptr, new_size);
  if (!ret)
  abort();
  return ret;
}

__attribute__((weak, export_name("canonical_abi_free")))
void canonical_abi_free(
void *ptr,
size_t size,
size_t align
) {
  free(ptr);
}
__attribute__((import_module("integers"), import_name("a1")))
void __wasm_import_integers_a1(int32_t);
void integers_a1(uint8_t x) {
  __wasm_import_integers_a1((int32_t) (x));
}
__attribute__((import_module("integers"), import_name("a2")))
void __wasm_import_integers_a2(int32_t);
void integers_a2(int8_t x) {
  __wasm_import_integers_a2((int32_t) (x));
}
