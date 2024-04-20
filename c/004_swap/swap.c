// This C program demonstrates the concept of swapping two variables using
// pointers.
#include <stdio.h>

// The swap function takes two pointers to integers as arguments. These pointers
// are used to directly manipulate the memory addresses of the variables they
// point to. Inside the swap function, a temporary variable t is declared and
// assigned the value of the integer pointed to by pa. Then, the value at the
// memory location pointed to by pa is replaced with the value at the memory
// location pointed to by pb. Finally, the value at the memory location pointed
// to by pb is replaced with the value stored in t. This effectively swaps the
// values of the variables that pa and pb point to.
void swap(int *pa, int *pb) {
    int t = *pa;
    *pa = *pb;
    *pb = t;
    return;
}

int main(void) {
    // In the main function, two integer variables a and b are declared and
    // initialized with the values 21 and 17 respectively.
    int a = 21;
    int b = 17;
    // The addresses of a and b are then passed to the swap function using the
    // address-of operator &. After the swap function is called, the values of a
    // and b are swapped.
    swap(&a, &b);
    // The printf function is then used to print the values of a and b to the
    // console, demonstrating that their values have indeed been swapped.
    printf("main: a = %d, b = %d\n", a, b);
    // The program then returns 0, signaling successful execution to the
    // operating system.
    return 0;
}
