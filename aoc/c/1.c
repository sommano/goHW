int accum = 0;
int change;
unsigned byte, bit;
char *bitfield;

if (!(bitfield = calloc(UINT_MAX / 8, 1)))
    err(1, "calloc");