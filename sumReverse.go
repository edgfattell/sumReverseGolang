/*
 * File: sumReverse.go
 * Author: Edward Fattell
 * Purpose: Takes user input for any number of integers on CLI, finds the
 *          reverse of said integer by listing its digits in reverse order, then
 *          adds each original integer to its reverse and prints the result to
 *          CLI.
 */
package main

import (
  "fmt"
  "math"
  "os"
  "io"
)

func main() {
  var n = 0;
  
  retVal , err := fmt.Scanf("%d", &n)

  for err == nil {
    if n > 0 {
      fmt.Printf("%d\n", n + revInt(n))
    } else {
      fmt.Fprintf(os.Stderr, "Error: input value %d is not positive\n", n)
    }
      retVal , err = fmt.Scanf("%d", &n)
  }

  if retVal == 0 && err != io.EOF {
    fmt.Fprintf(os.Stderr, "Error: Non-integer value in input\n");
  }
}

/*
 * revInt(n int) -- finds an integer whose digits are those of an integer n in
 * reverse order by finding the digit in n's 1's slot using mod, multiplying
 * that digit by 10 to the power of that digits "reverse" slot, adding that
 * digit to a seperate integer, and diving the original integer by 10, until the
 * whole number has been parsed. Then the result is returned.
 * Assumes n is not negative.
 */
func revInt(n int) int {
  var result = 0
  var index = int(math.Pow(10, math.Floor(math.Log10(float64(n)))))

  for i := n; i > 0; i /= 10 {
    var lastDig = i % 10 

    result += lastDig * index
    index /= 10 
  }
  return result
}
