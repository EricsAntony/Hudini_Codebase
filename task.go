//https://www.codewars.com/kata/56f69d9f9400f508fb000ba7

package kata
// monkeyCount takes the number n and returns the array of count.
func monkeyCount(n int) []int {
   // Your code here, happy coding!
   count:= make([]int,n)
   for i:=0; i<n; i++ {
     count[i] = i+1;
   }
  //returns the count;
  return count;
}


//https://www.codewars.com/kata/55685cd7ad70877c23000102

package kata

// MakeNegative accepts an integer and returns corresponding -value based on the input.
func MakeNegative(x int) int {
  if x == 0 {
    return 0
  } else if x > 0{
    return -x
  } else if x < 0 {
    return x;
  }
  return 0;
}


//https://www.codewars.com/kata/58ca658cc0d6401f2700045f

package kata

//FindMultiples return the multiples of the given integer upto the limit and returns the slice.
func FindMultiples(integer, limit int) []int {
  // Your code here!
  multiples := make([]int, limit/integer)
  index := 0;
  for i := 1; i <=limit; i++ {
    if i%integer == 0 {
      multiples[index] = i;
      index++;
    } 
  }
  return multiples;
}


//https://www.codewars.com/kata/5513795bd3fafb56c200049e

package kata

// CountBy function takes two integers and returns the slice that contain multiples of x upto the given n.
func CountBy(x, n int) []int {
  multiples := make([]int, n);
  for i := 0; i < n; i++ {
    multiples[i] = (i+1)*x; 
  }
  // returns the slice
  return multiples;
}

//https://www.codewars.com/kata/57a083a57cb1f31db7000028

package kata

import "math"

// PowersOfTwo takes n as input and returns the slice with exponents of two ranging from 0 to n
func PowersOfTwo(n int) []uint64 {
  // your code goes here
  powers := make([]uint64, n+1)
  for i := 0; i < n+1; i++ {
    powers[i] = uint64(math.Pow(2, float64(i)));
  }
  // return the powers
  return powers;
}

//https://www.codewars.com/kata/5bb904724c47249b10000131

package kata
import (
  "strings"
  "strconv"
)

func Points(games []string) int {
  // your code here!
  points := 0;
  x,y := 0,0;
  var str []string;
  for _,xi := range games {
    str = strings.Split(xi,":");
    x,_ = strconv.Atoi(str[0]);
    y,_ = strconv.Atoi(str[1]);
    
    if x > y {
      points += 3;
    } else if x < y {
      points += 0;
    } else if x == y {
      points += 1;
    }
  }
  return points
}


// https://www.codewars.com/kata/55f2b110f61eb01779000053

package kata
// GetSum takes two parameters and returns sum of all numbers between the inputs.
func GetSum(a, b int) int {
  sum := 0;
  max := b;
  min := a;
  if a > b {
    max = a;
    min = b;
  }
  if a == b {
    return a;
  }
  for i := min; i <= max; i++ {
    sum += i;
  }
  return sum
}

// https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

package kata
import (
"strings"
)

// FindShort takes a string and returns the length of the shortest word
func FindShort(s string) int {
  // your code
  var str []string;
  var shortest string;
  str = strings.Split(s," ");
  shortest = str[0];
  for i := 0; i < len(str); i++ {
    if len(str[i]) < len(shortest) {
      shortest = str[i];
    }
  }
  return len(shortest);  
}

// https://www.codewars.com/kata/598f76a44f613e0e0b000026

package kata

import (
  "strconv"
  "unicode"
)

// SumOfIntegersInString takes a string and returns the sum of integers present in the string.
func SumOfIntegersInString(strng string) int {
  sum := 0;
  numberInString := "";
  str := []rune(strng)
  
  // Loop iterate though the rune and checks if the character is a digit and calculates the sum.
  for i := 0; i < len(str); i++ {
    if unicode.IsDigit(str[i]) {
      numberInString += string(str[i]);
    } else {
      if numberInString != "" {
        convertedNumber,_ := strconv.Atoi(numberInString);
        sum += convertedNumber;
        numberInString = "";
      }
    }
  }

  convertedNumber,_ := strconv.Atoi(numberInString);
  sum += convertedNumber; 

  return sum;
}

//https://www.codewars.com/kata/5f8341f6d030dc002a69d7e4

package kata


func LeastLarger(a []int, i int) int {
  givenIndexElement := a[i];
  minElement := -9999;
  minElementIndex := -1;
  for j := 0; j < len(a); j++ {
    if a[j] > givenIndexElement {
      if minElement == -9999 {
        minElement = a[j];
        minElementIndex = j;
      } else if (a[j] < minElement) {
        minElement = a[j];
        minElementIndex = j;
      }
    }
  }
  
  return minElementIndex;
}


//https://www.codewars.com/kata/515e271a311df0350d00000f

package kata

import (
"fmt"
"math"
)

func SquareSum(numbers []int) int {
    // your code here
  squareSum := 0;
  for i := 0; i < len(numbers);  i++ {
    fmt.Printf("%d\n", numbers[i] ^2);
    squareSum += int(math.Pow(float64(numbers[i]),2));
  }
  return squareSum;
}


//https://www.codewars.com/kata/59342039eb450e39970000a6

package kata

func OddCount(n int) int{
  //your code here
  
  return n/2
}

//https://www.codewars.com/kata/580755730b5a77650500010c

package kata

func SortMyString(s string) string {
  var evenIndexedString string;
  var oddIndexedString string;
  
  for i, xi := range s {
    if i % 2 == 0 {
      evenIndexedString += string(xi);
    } else {
      oddIndexedString += string(xi); 
    }
  }
  return evenIndexedString+" "+oddIndexedString;
}
