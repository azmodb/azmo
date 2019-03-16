#ifndef __INTERNAL_BINARY_H
#define __INTERNAL_BINARY_H

#include <string.h>
#include <stdint.h>

/* binary.c */
unsigned char* guint64(unsigned char* p, unsigned char* ep, uint64_t* v);
unsigned char* guint32(unsigned char* p, unsigned char* ep, uint32_t* v);
unsigned char* guint16(unsigned char* p, unsigned char* ep, uint16_t* v);
unsigned char* guint8(unsigned char* p, unsigned char* ep, uint8_t* v);
unsigned char* gstring(unsigned char* p, unsigned char* ep, char** v);
unsigned char* gdata(unsigned char* p, unsigned char* ep, char** v);

unsigned char* puint64(unsigned char* p, unsigned char* ep, uint64_t v);
unsigned char* puint32(unsigned char* p, unsigned char* ep, uint32_t v);
unsigned char* puint16(unsigned char* p, unsigned char* ep, uint16_t v);
unsigned char* puint8(unsigned char* p, unsigned char* ep, uint8_t v);
unsigned char* pstring(unsigned char* p, unsigned char* ep, char* v);
unsigned char* pdata(unsigned char* p, unsigned char* ep, char* v);

#endif /* __INTERNAL_BINARY_H */
