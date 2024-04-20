// This C program demonstrates the use of a static variable within a function.
#include <stdio.h>

void increment(void) {
    // The increment function contains a static unsigned integer variable
    // counter which is initialized to 0. The keyword static means that the
    // counter variable retains its value between function calls. This is
    // different from regular local variables which are destroyed and recreated
    // each time a function is called.
    static unsigned int counter = 0;  // static duration, block scope
    counter++;
    printf("%d\n", counter);
}

int main(void) {
    for (int i = 0; i < 10; i++) {
        increment();
    }
    return 0;
}
