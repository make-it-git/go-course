#include <stdio.h>

int main () {
   FILE *fp;
   char str[] = "example";

   fp = fopen("/tmp/go-course.txt" , "w");
   size_t written;
   // size_t fwrite(const void *ptr, size_t size, size_t count, FILE *stream);
   written = fwrite(str, 1 , sizeof(str) , fp);

   printf("Written to file %zu\n", written);

   fclose(fp);

   return 0;
}