#include <stdio.h>
#include <stdlib.h>

int scoreOfString(char* s) {
  int i, sum = 0;
  for (i = 1; s[i]; ++i) {
    sum += abs(s[i-1] - s[i]);
  }
  return sum;
}

int main() {
  int ans = scoreOfString("hello");

  printf("ans: %d\n", ans);

  return 0;
}
