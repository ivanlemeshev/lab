// This C code snippet defines a structure S and demonstrates the concept of
// memory alignment in C.

// This code demonstrates the importance of understanding and respecting the memory alignment requirements of your data types in C. Misaligned access can lead to crashes, incorrect behavior, or significant performance penalties on some architectures.

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
    // In the main function, two buffers are declared: bad_buff and good_buff.
    // Both are arrays of unsigned char with size equal to the size of struct S.
    // The difference between them is that good_buff is declared with the
    // _Alignas keyword, which enforces the alignment requirements of struct S
    // on good_buff. This means that good_buff will start at a memory address
    // that is suitable for a struct S object.
    unsigned char bad_buff[sizeof(struct S)];
    _Alignas(struct S) unsigned char good_buff[sizeof(struct S)];

    // Then, two pointers bad_s_ptr and good_s_ptr are declared. Both are
    // pointers to struct S. They are initialized by casting the addresses of
    // bad_buff and good_buff respectively to struct S*.

    // Yhe bad_s_ptr line is wrong pointer alignment. This is because bad_buff
    // was not guaranteed to be properly aligned for a struct S object, so
    // casting its address to struct S* could lead to undefined behavior if the
    // pointer is dereferenced.
    struct S *bad_s_ptr = (struct S *)bad_buff;

    // On the other hand, good_s_ptr is correct pointer alignment because
    // good_buff was declared with _Alignas(struct S), ensuring that it is
    // properly aligned for a struct S object. Therefore, casting its address to
    // struct S* is safe.
    struct S *good_s_ptr = (struct S *)good_buff;
}
