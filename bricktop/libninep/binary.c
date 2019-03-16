#include "binary.h"

unsigned char*
guint64(unsigned char* p, unsigned char* ep, uint64_t* v)
{
	if(p == NULL || p + 8 > ep)
		return NULL;

	*v = (uint64_t)p[0] | ((uint64_t)p[1] << 8) | ((uint64_t)p[2] << 16) |
	     ((uint64_t)p[3] << 24) | ((uint64_t)p[4] << 32) |
	     ((uint64_t)p[5] << 40) | ((uint64_t)p[6] << 48) |
	     ((uint64_t)p[7] << 56);
	p += 8;
	return p;
}

unsigned char*
guint32(unsigned char* p, unsigned char* ep, uint32_t* v)
{
	if(p == NULL || p + 4 > ep)
		return NULL;

	*v = p[0] | p[1] << 8 | p[2] << 16 | p[3] << 24;
	p += 4;
	return p;
}

unsigned char*
guint16(unsigned char* p, unsigned char* ep, uint16_t* v)
{
	if(p == NULL || p + 2 > ep)
		return NULL;

	*v = p[0] | p[1] << 8;
	p += 2;
	return p;
}

unsigned char*
guint8(unsigned char* p, unsigned char* ep, uint8_t* v)
{
	if(p == NULL || p + 1 > ep)
		return NULL;

	*v = p[0];
	p += 1;
	return p;
}

unsigned char*
gstring(unsigned char* p, unsigned char* ep, char** v)
{
	if(p == NULL || p + 2 > ep)
		return NULL;

	uint16_t size = p[0] | p[1] << 8;
	p += 1;
	if(p + size + 1 > ep)
		return NULL;
	memmove(p, p + 1, size);
	p[size] = '\0';
	*v      = (char*)p;
	p += size + 1;
	return p;
}

unsigned char*
gdata(unsigned char* p, unsigned char* ep, char** v)
{
	if(p == NULL || p + 4 > ep)
		return NULL;

	uint32_t size;
	size = p[0] | p[1] << 8 | p[2] << 16 | p[3] << 24;
	p += 3;
	if(p + size + 3 > ep)
		return NULL;
	memmove(p, p + 3, size);
	p[size] = '\0';
	*v      = (char*)p;
	p += size + 3;
	return p;
}

unsigned char*
puint64(unsigned char* p, unsigned char* ep, uint64_t v)
{
	if(p == NULL || p + 8 > ep)
		return NULL;

	p[0] = v;
	p[1] = v >> 8;
	p[2] = v >> 16;
	p[3] = v >> 24;
	p[4] = v >> 32;
	p[5] = v >> 40;
	p[6] = v >> 48;
	p[7] = v >> 56;
	p += 8;
	return p;
}

unsigned char*
puint32(unsigned char* p, unsigned char* ep, uint32_t v)
{
	if(p == NULL || p + 4 > ep)
		return NULL;

	p[0] = v;
	p[1] = v >> 8;
	p[2] = v >> 16;
	p[3] = v >> 24;
	p += 4;
	return p;
}

unsigned char*
puint16(unsigned char* p, unsigned char* ep, uint16_t v)
{
	if(p == NULL || p + 2 > ep)
		return NULL;

	p[0] = v;
	p[1] = v >> 8;
	p += 2;
	return p;
}

unsigned char*
puint8(unsigned char* p, unsigned char* ep, uint8_t v)
{
	if(p == NULL || p + 1 > ep)
		return NULL;

	p[0] = v;
	p += 1;
	return p;
}

unsigned char*
pstring(unsigned char* p, unsigned char* ep, char* v)
{
	if(v == NULL)
		return puint16(p, ep, 0);

	size_t n = strlen(v);
	if(p == NULL || p + 2 + n > ep)
		return NULL;

	p = puint16(p, ep, n);
	if(p == NULL)
		return NULL;
	memmove(p, v, n);
	p += n;
	return p;
}

unsigned char*
pdata(unsigned char* p, unsigned char* ep, char* v)
{
	if(v == NULL)
		return puint32(p, ep, 0);

	size_t n = strlen(v);
	if(p == NULL || p + 4 + n > ep)
		return NULL;

	p = puint32(p, ep, n);
	if(p == NULL)
		return NULL;
	memmove(p, v, n);
	p += n;
	return p;
}
