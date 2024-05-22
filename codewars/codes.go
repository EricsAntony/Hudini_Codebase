// https://www.codewars.com/kata/586f6741c66d18c22800010a/train/go

package kata


func SequenceSum(start, end, step int) int {
  sum := 0
  if start > end {
    return 0
  }
  for i:=start; i<=end; i+=step{
    sum += i
  }
  return sum
}

// https://www.codewars.com/kata/59f38b033640ce9fc700015b/train/go

package kata 

func Solve(arr []int) int {
  sum := 0
  for i, value := range arr{
    if isPrime(i) {
      sum += value;
    }
  }
  return sum
}

func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

//https://www.codewars.com/kata/554e4a2f232cdd87d9000038/train/go

package kata

func DNAStrand(dna string) string {
  // your code here
  str := "";
  for _, value := range dna {
    if value == 'A' {
      str += "T"
    } else if value == 'T'{
      str += "A"
    } else if value == 'C'{
      str += "G"
    } else if value == 'G' {
      str += "C"
    } else {
      str += string(value)
    }
  }
  return str;
}

// https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go

package kata
import("fmt")
func FindOdd(seq []int) int {
  count := make(map[int]int)
  
  for _, value := range seq {
    _, ok := count[value]
    if ok {
      count[value]++
    } else {
      count[value]=1
    }
  }
  fmt.Print(count)
  for key, _ := range count{
    if count[key]%2 != 0 {
      return key
    }
  }
  return 0
}

// https://www.codewars.com/kata/59e49b2afc3c494d5d00002a/train/go

package kata
import "strings"
func SortVowels(s string) string {
	str := ""
  for i := 0; i < len(s); i++ {
    if string(s[i]) == "A" || string(s[i]) == "E" || string(s[i]) == "I" || string(s[i]) == "O" ||
    string(s[i]) == "U" {
      str += "|"+string(s[i])+"\n"
    } else if string(s[i]) == "a" || string(s[i]) == "e" || string(s[i]) == "i" || string(s[i]) == "o" ||
    string(s[i]) == "u" {
      str += "|"+string(s[i])+"\n"
    } else {
      str += string(s[i])+"|\n"
    }
  }
  return strings.TrimSuffix(str, "\n");
}

// https://www.codewars.com/kata/5679aa472b8f57fb8c000047/train/go

package kata
import ("fmt")

func FindEvenIndex(arr []int) int {
  sumLeft := 0
  sumRight := 0
  for i := 0; i < len(arr); i++ {
    sumLeft = calculateSum(0,i-1,arr)
	fmt.Printf("left %d \n", sumLeft);

    sumRight = calculateSum(i+1, len(arr)-1,arr)
	fmt.Printf("right %d \n", sumRight);
	fmt.Printf("\n");

    if sumLeft == sumRight {
      return i
    }
  }
  return -1
}

func calculateSum(start, end int, arr []int) int {
  
  sum := 0
  for i := start; i <= end; i++ {
    sum += arr[i];
 }
  return sum
}

// https://www.codewars.com/kata/592a6ad46d6c5a62b600003f

package kata

import (
    "strings"
)

var cipherMap = map[rune]rune{
    'G': 'A', 'A': 'G', 'g': 'a', 'a': 'g',
    'D': 'E', 'E': 'D', 'd': 'e', 'e': 'd',
    'R': 'Y', 'Y': 'R', 'r': 'y', 'y': 'r',
    'P': 'O', 'O': 'P', 'p': 'o', 'o': 'p',
    'L': 'U', 'U': 'L', 'l': 'u', 'u': 'l',
    'K': 'I', 'I': 'K', 'k': 'i', 'i': 'k',
}

func Encode(input string) string {
    return transform(input)
}

func Decode(input string) string {
    return transform(input)
}

func transform(input string) string {
    var result strings.Builder
    for _, char := range input {
        if substitute, ok := cipherMap[char]; ok {
            result.WriteRune(substitute)
        } else {
            result.WriteRune(char)
        }
    }
    return result.String()
}