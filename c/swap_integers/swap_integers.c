#include <stdio.h>
#include <stdlib.h>

void swap(int *pa, int *pb) {
  int t = *pa;
  *pa = *pb;
  *pb = t;
}

int main() {
  int a = 5;
  int b = 10;
  printf("main: a = %d, b = %d\n", a, b);
  swap(&a, &b);
  printf("main: a = %d, b = %d\n", a, b);
  return EXIT_SUCCESS;
}
