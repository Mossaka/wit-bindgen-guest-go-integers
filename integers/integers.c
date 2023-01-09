#include "integers.h"


__attribute__((import_module("imports"), import_name("a1")))
void __wasm_import_imports_a1(int32_t);

__attribute__((import_module("imports"), import_name("a2")))
void __wasm_import_imports_a2(int32_t);

__attribute__((weak, export_name("cabi_realloc")))
void *cabi_realloc(void *ptr, size_t orig_size, size_t org_align, size_t new_size) {
  void *ret = realloc(ptr, new_size);
  if (!ret) abort();
  return ret;
}

// Component Adapters

void imports_a1(uint8_t x) {
  __wasm_import_imports_a1((int32_t) (x));
}

void imports_a2(int8_t x) {
  __wasm_import_imports_a2((int32_t) (x));
}

__attribute__((export_name("res")))
int32_t __wasm_export_integers_res(void) {
  uint8_t ret = integers_res();
  return (int32_t) (ret);
}

extern void __component_type_object_force_link_integers(void);
void __component_type_object_force_link_integers_public_use_in_this_compilation_unit(void) {
  __component_type_object_force_link_integers();
}
