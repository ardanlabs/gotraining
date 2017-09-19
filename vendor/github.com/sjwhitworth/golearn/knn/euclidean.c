// #cgo CFLAGS: -Og -march=native -ffast-math

#include <stdio.h>
#include <string.h>
#include "knn.h"

/* Works out the Euclidean distance (not square-rooted) for a given
 * AttributeGroup */
void euclidean_distance (
        struct dist *out, /* Output distance vector, needs to be initially zero */
        int max_row,      /* Size of the output vector */
        int max_col,      /* Number of columns */
        int row,          /* Current row */
        double *train,    /* Pointer to first element of training AttributeGroup */
        double *pred      /* Pointer to first element of equivalent prediction AttributeGroup */
    )
{
    int i, j;
    for (i = 0; i < max_row; i++) {
        out[i].p = i;
        for (j = 0; j < max_col; j++) {
            double tmp;
            tmp  = *(pred  + row * max_col   + j);
            tmp -= *(train + i   * max_col   + j);
            tmp *= tmp; /* Square */
            out[i].dist += tmp;
        }
    }
}

