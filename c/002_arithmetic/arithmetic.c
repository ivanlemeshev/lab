#include <stdio.h>
#include <stdlib.h>

int main(void) {
  int a = 17;
  int b = 5;

  printf("a = %d, b = %d\n", a, b);
  printf("Sum: %d\n", a + b);
  printf("Difference: %d\n", a - b);
  printf("Product: %d\n", a * b);
  printf("Quotient: %d\n", a / b);
  printf("Remainde: %d\n", a % b);

  int x = 17 / 5;
  double y = 17 / 5;
  double z = 17.0 / 5;

  printf("%d\n", x);
  printf("%f\n", y);
  printf("%f\n", z);

  return EXIT_SUCCESS;
}
