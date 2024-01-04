#include <stdio.h>
#include <stdlib.h>
#include "../../binaries/macos-arm64/poseidon.h"

int main() {
    // input arguments must be field element represented as \0 terminated string
    const char *x = "123456789";
    const char *y = "9876543210";

    // result memory is managed by C
    char *result = Hash(x, y);
    printf("poseidon(%s, %s) = %s", x, y, result);

    // free result memory
    free(result);
}