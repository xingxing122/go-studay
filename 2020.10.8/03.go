package main

import "fmt"

func main() {
   a,b,c := 2,3,5
   for a := 7; a < 8; a++ {
   	b := 11
   	c = 13
   	fmt.Println("inner: a-%d b- %d c-%d\n", a,b,c)

   }
   fmt.Println("outer: a-d%  b-%d  c-d%\n",a,b,c)

}
