// This C code snippet defines a structure S and demonstrates the concept of
// memory alignment in C.

// The line includes the stdalign.h header file. This header file, introduced in
// C11, defines two macros: alignas and alignof. The alignas macro allows you to
// specify the alignment of variables or types, while the alignof macro allows
// you to retrieve the alignment requirement of a type. These macros are useful
// for low-level programming tasks that require precise control over the memory
// layout of your data.
#include <stdalign.h>

// The line includes the stdio.h header file. This is one of the most commonly
// used header files in C. It defines several types, macros, and functions for
// performing input and output, including printf, scanf, FILE, EOF, and many
// others. This header file is typically included in any C program that needs to
// interact with the standard input, standard output, or files.
#include <stdio.h>

// This code demonstrates the importance of understanding and respecting the
// memory alignment requirements of your data types in C. Misaligned access can
// lead to crashes, incorrect behavior, or significant performance penalties on
// some architectures.

// The structure S contains three members: an integer i, a double d, and a
// character c. The size of this structure will be determined by the size of
// these three members and the alignment requirements of the system it's
// compiled on.
struct S {
    int i;
    double d;
    char c;
};

int main(void) {
    // The sizeof operator in C returns the size (in bytes) of the object or
    // type that follows it. It's a compile-time operator, meaning the sizes are
    // determined when the program is compiled, not when it's run.
    printf("sizeof(int) = %zu byte(s)\n", sizeof(int));
    printf("sizeof(double) = %zu byte(s)\n", sizeof(double));
    printf("sizeof(char) = %zu byte(s)\n", sizeof(char));
    printf("sizeof(struct S) = %zu byte(s)\n", sizeof(struct S));

    // In the main function, two buffers are declared: bad_buff and good_buff.
    // Both are arrays of unsigned char with size equal to the size of struct S.
    // The difference between them is that good_buff is declared with the
    // _Alignas keyword, which enforces the alignment requirements of struct S
    // on good_buff. This means that good_buff will start at a memory address
    // that is suitable for a struct S object.
    unsigned char bad_buff[sizeof(struct S)];
    printf("sizeof(unsigned char bad_buff) = %zu byte(s)\n", sizeof(bad_buff));
    _Alignas(struct S) unsigned char good_buff[sizeof(struct S)];
    printf("sizeof(unsigned char good_buff) = %zu byte(s)\n",
           sizeof(good_buff));

    // Then, two pointers bad_s_ptr and good_s_ptr are declared. Both are
    // pointers to struct S. They are initialized by casting the addresses of
    // bad_buff and good_buff respectively to struct S*.

    // The bad_s_ptr line is wrong pointer alignment. This is because bad_buff
    // was not guaranteed to be properly aligned for a struct S object, so
    // casting its address to struct S* could lead to undefined behavior if the
    // pointer is dereferenced.
    struct S *bad_s_ptr = (struct S *)bad_buff;
    printf("sizeof(*bad_s_ptr) = %zu byte(s)\n", sizeof(*bad_s_ptr));

    // On the other hand, good_s_ptr is correct pointer alignment because
    // good_buff was declared with _Alignas(struct S), ensuring that it is
    // properly aligned for a struct S object. Therefore, casting its address to
    // struct S* is safe.
    struct S *good_s_ptr = (struct S *)good_buff;
    printf("sizeof(*good_s_ptr) = %zu byte(s)\n", sizeof(*good_s_ptr));
}
