#ifndef _H_FUNCS
#define _H_FUNCS

#include <stdint.h>

struct dist {
    float dist;
    uint32_t p;
};

/* Works out the Euclidean distance (not square-rooted) for a given
 * AttributeGroup */
void euclidean_distance (
        struct dist *out, /* Output distance vector, needs to be initially zero */
        int max_row,      /* Size of the output vector */
        int max_col,      /* Number of columns */
        int row,          /* Current prediction row */
        double *train,    /* Pointer to first element of training AttributeGroup */
        double *pred      /* Pointer to first element of equivalent prediction AttributeGroup */
); 
#endif
