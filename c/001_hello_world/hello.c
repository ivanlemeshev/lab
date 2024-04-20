// This is a preprocessor directive that includes the standard input/output
// library in the program. This library contains functions for input and output
// operations. In this case, it's needed for the puts() function used later in
// the program.
#include <stdio.h>

// This is another preprocessor directive that includes the standard library in
// the program. This library contains general purpose functions, including
// memory management, process control, conversions and others. In this case,
// it's needed for the EXIT_SUCCESS macro used later in the program.
#include <stdlib.h>

// This is the main function of the program. In C, execution of the program
// starts with the main function. The void keyword indicates that this function
// takes no arguments.
int main(void) {
    // This line calls the puts() function, which is a standard library function
    // that writes a string to the console, followed by a newline character.
    // Here, it's used to print "Hello, World!".
    puts("Hello, World!");
    // This line indicates that the program has finished successfully.
    // EXIT_SUCCESS is a macro (defined in stdlib.h) that represents a
    // successful program termination status. In most systems, it's defined as
    // 0. If a C program finishes without any return statement in main, it
    // returns a status of EXIT_SUCCESS by default.
    return EXIT_SUCCESS;
}
