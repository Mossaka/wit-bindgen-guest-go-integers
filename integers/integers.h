#ifndef __BINDINGS_INTEGERS_H
#define __BINDINGS_INTEGERS_H
#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

// Imported Functions from `imports`
void imports_a1(uint8_t x);
void imports_a2(int8_t x);

// Exported Functions from `integers`
uint8_t integers_res(void);

#ifdef __cplusplus
}
#endif
#endif
