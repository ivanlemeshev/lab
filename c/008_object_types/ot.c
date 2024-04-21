// -----------------------------------------------------------------------------
// This C code snippet demonstrates the use of boolean variables in C.
// -----------------------------------------------------------------------------
#include <stdio.h>
#include <stdlib.h>

// The line includes the stdbool.h header file. This header file, introduced in
// C99, defines four macros: bool, true, false, and
// __bool_true_false_are_defined. The bool macro expands to _Bool, which is the
// actual keyword for boolean data type in C. The true and false macros expand
// to 1 and 0 respectively.
#include <stdbool.h>

// The line declares a boolean variable flag1 and initializes it to 0. In C,
// _Bool is a built-in data type that can hold one of two values: 0 for false,
// and 1 for true.
_Bool flag1 = 0;

// The line declares a boolean variable flag2 and initializes it to false. Here,
// bool is a macro that expands to _Bool, and false is a macro that expands to
// 0. So, this line is equivalent to _Bool flag2 = 0;.
bool flag2 = false;

// In both cases, the boolean variables are initialized to false (or 0). The use
// of bool, true, and false macros from stdbool.h can make your C code easier to
// read and understand, because it allows you to use boolean logic in a way that
// is similar to other high-level languages.

// -----------------------------------------------------------------------------
// This C code snippet demonstrates the use of different character types in C.
// -----------------------------------------------------------------------------

// This code demonstrates the use of different character types in C, including
// both narrow characters (char, unsigned char, signed char) and wide characters
// (wchar_t). The choice between these types depends on the requirements of your
// program. For example, if you need to handle Unicode characters that can't be
// represented in a single byte, you would use wchar_t.

// The line includes the wchar.h header file. This header file defines several
// types and functions for working with wide characters, which are characters
// that can have a width of more than one byte.
#include <wchar.h>

// The line declares a character variable c and initializes it with the
// character 'a'. In C, char is a built-in data type that can hold a single
// byte. By default, it can be either signed or unsigned, depending on the
// compiler and the platform.
char c = 'a';

// The line declares an unsigned character variable uc and initializes it with
// the character 'b'. An unsigned char is a character data type that can hold a
// single byte, and its range of values is from 0 to 255.
unsigned char uc = 'b';
// The line declares a signed character variable sc and initializes it with the
// character 'c'. A signed char is a character data type that can hold a single
// byte, and its range of values is from -128 to 127.

signed char sc = 'c';
// The line declares a wide character variable wc and initializes it with the
// wide character 'Ф'. The wchar_t is a type defined in the wchar.h header file.
// It's used to store wide characters. The L before the character literal 'Ф'
// indicates that it's a wide character literal.

wchar_t wc = L'Ф';

// -----------------------------------------------------------------------------
// This C code snippet declares five different integral types: signed char,
// short int, int, long int, and long long int.
// -----------------------------------------------------------------------------

// The line declares a variable named itsc of type signed char. A signed char is
// a character data type that can hold a single byte, and its range of values is
// from -128 to 127.
signed char itsc;

// The line declares a variable named itsi of type short int. A short int is an
// integer type that is usually used to save memory in large arrays. It has a
// smaller range than int, typically from -32768 to 32767.
short int itsi;

// The line declares a variable named iti of type int. An int is the most
// commonly used integer type. Its size and range can vary depending on the
// system, but it's usually at least 16 bits, with a range from -32768 to 32767
// on a 16-bit system. It is in the range of -2147483648 to 2147483647 on a
// 32-bit and 64-bit systems.
int iti;

// The line declares a variable named itli of type long int. A long int is an
// integer type that has a larger range than int. Its size and range can vary
// depending on the system, but it's usually at least 32 bits, with a range from
// -2147483648 to 2147483647 on a 32-bit system. It is in the range of
// -9223372036854775808 to 9223372036854775807 on a 64-bit system.
long int itli;

// The line declares a variable named itlli of type long long int. A long long
// int is an integer type that has an even larger range than long int. Its size
// is usually 64 bits, with a range from -9223372036854775808 to
// 9223372036854775807.
long long int itlli;

int main(void) {
    printf("sizeof(_Bool flag1) = %zu byte(s)\n", sizeof(flag1));
    printf("sizeof(bool flag2) = %zu byte(s)\n", sizeof(flag2));

    printf("sizeof(char c) = %zu byte(s)\n", sizeof(c));
    printf("sizeof(unsigned char uc) = %zu byte(s)\n", sizeof(uc));
    printf("sizeof(signed char sc) = %zu byte(s)\n", sizeof(sc));
    printf("sizeof(wchar_t wc) = %zu byte(s)\n", sizeof(wc));

    printf("sizeof(signed char itsc) = %zu byte(s)\n", sizeof(itsc));
    printf("sizeof(short int itsi) = %zu byte(s)\n", sizeof(itsi));
    printf("sizeof(int iti) = %zu byte(s)\n", sizeof(iti));
    printf("sizeof(long int itli) = %zu byte(s)\n", sizeof(itli));
    printf("sizeof(long long int itlli) = %zu byte(s)\n", sizeof(itlli));

    return EXIT_SUCCESS;
}
