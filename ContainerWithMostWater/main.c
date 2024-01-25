#include <stdio.h>

int min(int a, int b) {
  if (a < b)
    return a;
  return b;
}

int maxArea(int *height, int heightSize) {
  int l = 0, r = heightSize - 1;
  int area = 0, max = 0;

  while (l < r) {
    area = min(height[l], height[r]) * (r - l);

    if (area > max)
      max = area;

    if (height[l] < height[r])
      l++;
    else
      r--;
  }

  return max;
}

int main() {
  int numbers[] = {7, 3, 8, 4, 5, 2, 6, 8, 1};
  int max = maxArea(numbers, 9);
  printf("%d", max);
}
