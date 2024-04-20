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
    // This line calls the printf() function, which is a standard library
    // function used for formatted output. The first argument is a format string
    // that specifies how subsequent arguments are converted for output. In this
    // case, %s is a placeholder for a string, and \n is a newline character.
    // The second argument, "Hello, world!", is the string that replaces the %s
    // placeholder in the format string. The condition checks if the result of
    // the printf() function is equal to EOF. EOF is a macro that represents the
    // end-of-file condition. In the context of the printf() function, it
    // returns the number of characters transmitted, or EOF if an output or
    // encoding error occurred.
    if (printf("%s\n", "Hello, world!") == EOF) {
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
