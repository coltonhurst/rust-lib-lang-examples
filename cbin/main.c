/*
    Compile this file with:
    
    gcc main.c -L../rustlib/target/debug -lrustlib -o main
*/

#include <stdio.h>

extern int add(int left, int right);

int main() {
    int result = add(5, 3);
    printf("The sum is: %d\n", result);

    return 0;
}
