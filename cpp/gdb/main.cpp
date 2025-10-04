#include <cstdio>

int do_something(int x) {
  if (x > 0) {
    return 1;
  }
  if (x == 0) {
    return 0;
  } else {
    return -1;
  }
}

int main() {
  int num1 = 42;
  int result1 = do_something(num1);
  int num2 = 0;
  int result2 = do_something(num2);
  int num3 = -32767;
  int result3 = do_something(num3);
  printf("Num1: %d, Step: %d\n", num1, result1);
  printf("Num2: %d, Step: %d\n", num2, result2);
  printf("Num3: %d, Step: %d\n", num3, result3);
  return 0;
}
