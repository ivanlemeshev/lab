#include <stdio.h>
#include <stdlib.h>

#include <stdbool.h>

_Bool flag1 = 0;

bool flag2 = false;

#include <wchar.h>

char c = 'a';
unsigned char uc = 'b';
signed char sc = 'c';
wchar_t wc = L'Ф';
signed char itsc;
short int itsi;
int iti;
long int itli;
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
