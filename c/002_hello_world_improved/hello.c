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
    // Here, it's used to print "Hello, World!". The condition checks if the
    // result of the puts() function is equal to EOF. EOF is a macro that
    // represents the end-of-file condition. In the context of the puts()
    // function, it returns EOF if an error occurs.
    if (puts("Hello, world!") == EOF) {
        //  If the condition in the if statement is true (i.e., an error
        //  occurred), this line is executed. EXIT_FAILURE is a macro (defined
        //  in stdlib.h) that represents an unsuccessful program termination
        //  status. In most systems, it's defined as a non-zero value. This line
        //  causes the program to terminate and return a status indicating that
        //  an error occurred.
        return EXIT_FAILURE;
    }

    // This line indicates that the program has finished successfully.
    // EXIT_SUCCESS is a macro (defined in stdlib.h) that represents a
    // successful program termination status. In most systems, it's defined as
    // 0. If a C program finishes without any return statement in main, it
    // returns a status of EXIT_SUCCESS by default.
    return EXIT_SUCCESS;
}
