#include <stdalign.h>
#include <stdio.h>

struct S {
  int i;
  double d;
  char c;
};

int main(void) {
  printf("sizeof(int) = %zu byte(s)\n", sizeof(int));
  printf("sizeof(double) = %zu byte(s)\n", sizeof(double));
  printf("sizeof(char) = %zu byte(s)\n", sizeof(char));
  printf("sizeof(struct S) = %zu byte(s)\n", sizeof(struct S));

  unsigned char bad_buff[sizeof(struct S)];
  printf("sizeof(unsigned char bad_buff) = %zu byte(s)\n", sizeof(bad_buff));
  _Alignas(struct S) unsigned char good_buff[sizeof(struct S)];
  printf("sizeof(unsigned char good_buff) = %zu byte(s)\n", sizeof(good_buff));

  struct S *bad_s_ptr = (struct S *)bad_buff;
  printf("sizeof(*bad_s_ptr) = %zu byte(s)\n", sizeof(*bad_s_ptr));

  struct S *good_s_ptr = (struct S *)good_buff;
  printf("sizeof(*good_s_ptr) = %zu byte(s)\n", sizeof(*good_s_ptr));
}
