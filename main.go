package main

/*
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void linear_regression(const char *filename) {
    FILE *file = fopen(filename, "r");
    if (!file) {
        printf("error reading file\n");
        return;
    }

    double *y = NULL;
    size_t size = 0;
    double value;

    while (fscanf(file, "%lf", &value) == 1) {
        y = realloc(y, (size + 1) * sizeof(double));
        if (!y) {
            printf("Memory allocation failed\n");
            fclose(file);
            return;
        }
        y[size++] = value;
    }
    fclose(file);

    double N = (double)size;
    if (N == 0) {
        printf("no data\n");
        free(y);
        return;
    }

    double sumX = 0.0, sumY = 0.0, sumXY = 0.0, sumX2 = 0.0, sumY2 = 0.0;

    for (size_t x = 0; x < size; x++) {
        sumX += (double)x;
        sumY += y[x];
        sumXY += (double)x * y[x];
        sumX2 += (double)x * x;
        sumY2 += y[x] * y[x];
    }

    double m = (N * sumXY - sumX * sumY) / (N * sumX2 - sumX * sumX);
    double b = (sumY - m * sumX) / N;

    double r = (N * sumXY - sumX * sumY) / sqrt((N * sumX2 - sumX * sumX) * (N * sumY2 - sumY * sumY));

    printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b);
    printf("Pearson Correlation Coefficient: %.10f\n", r);

    free(y);
}
*/
import "C"
import (
    "os"
    "unsafe"
)

func main() {
    if len(os.Args) < 2 {
        println("check args")
        return
    }

    cFilename := C.CString(os.Args[1])
    defer C.free(unsafe.Pointer(cFilename)) 
    C.linear_regression(cFilename)
}
