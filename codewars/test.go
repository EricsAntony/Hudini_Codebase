package main
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
      return arr[i]
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

func main() {
	FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1})
}
