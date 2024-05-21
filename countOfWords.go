package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main (){
	var text string;
	countOfWords := make(map[string]int)
	count := 0;
	fmt.Println("Enter the text: ");
	reader := bufio.NewReader(os.Stdin)
	text, _ = reader.ReadString('\n')
	words := strings.Split(strings.TrimSpace(text), " ");
	for i := 0; i < len(words); i++ {
		//fmt.Print(words[i]);
		_, exists := countOfWords[words[i]]
		if !exists{
			count = 1;
			for j := i+1; j<len(words); j++ {
				if words[i] == words[j]{
						count++
				}
			}
			countOfWords[words[i]] = count;
		}
		
	}

	for key, value := range countOfWords{
		fmt.Printf("%s : %d\n", key, value);
	}

}