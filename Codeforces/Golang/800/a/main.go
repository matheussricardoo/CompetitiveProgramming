package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	s := make([]float64, len(messages))
	count := 0.0
	for i:=0; i < len(messages);i++ {
		count += float64(len(messages[i])) * 0.01
		s = append(s, count)
	}

	return s 
}

func main () {
	strings := []string{"I don't want to be here anymore"}
	fmt.Println(getMessageCosts(strings))
	
}