#include <stdlib.h>
#include <integers-export.h>

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
__attribute__((export_name("res")))
int32_t __wasm_export_integers_export_res(void) {
  uint8_t ret = integers_export_res();
  return (int32_t) (ret);
}
